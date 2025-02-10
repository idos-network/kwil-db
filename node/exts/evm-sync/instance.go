package evmsync

import (
	"context"
	"fmt"
	"sync"

	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/kwilteam/kwil-db/common"
	"github.com/kwilteam/kwil-db/extensions/listeners"
	"github.com/kwilteam/kwil-db/node/exts/evm-sync/chains"
	orderedsync "github.com/kwilteam/kwil-db/node/exts/ordered-sync"
	"github.com/kwilteam/kwil-db/node/types/sql"
)

// RegisterEventResolution registers a resolution function for the EVM event listener.
// It should be called in an init function.
func RegisterEventResolution(name string, resolve ResolveFunc) {
	_, ok := registeredResolutions[name]
	if ok {
		panic(fmt.Sprintf("resolution with name %s already registered", name))
	}

	registeredResolutions[name] = resolve

	orderedsync.RegisterResolveFunc(name, func(ctx context.Context, app *common.App, block *common.BlockContext, res *orderedsync.ResolutionMessage) error {
		logs, err := deserializeLogs(res.Data)
		if err != nil {
			return err
		}

		return resolve(ctx, app, block, res.Topic, logs)
	})
}

var (
	// EventSyncer is the singleton that is responsible for syncing events from Ethereum.
	EventSyncer = &globalListenerManager{
		listeners: make(map[string]*listenerInfo),
	}

	// resolutions is a map of resolution functions
	registeredResolutions = make(map[string]ResolveFunc)
)

// this file contains a thread-safe in-memory cache for the chains that the network cares about.

type globalListenerManager struct {
	// mu protects all fields in this struct
	mu sync.Mutex
	// listeners is a set of listeners
	listeners map[string]*listenerInfo
	// shouldListen is a flag that is set to true when the node should have
	// its listeners running.
	shouldListen bool

	/*
		THE BELOW FIELDS ARE ONLY SET WHEN shouldListen IS TRUE
	*/

	// runningContext is the context that the listeners are running in.
	runningContext    context.Context
	runningService    *common.Service
	runningEventStore listeners.EventStore
	runningSyncConf   *syncConfig
}

// EVMEventListenerConfig is the configuration for an EVM event listener.
type EVMEventListenerConfig struct {
	// UniqueName is a unique name for the listener.
	// It MUST be unique from all other listeners.
	UniqueName string
	// Chain is the chain that the listener is listening to.
	Chain chains.Chain
	// GetLogs is a function that queries logs to be synced from the chain.
	GetLogs GetBlockLogsFunc
	// Resolve is the function that will be called the Kwil network
	// has confirmed events from Ethereum.
	Resolve string
}

// RegisterListener registers a new listener.
// It should be called when a node starts up (e.g. on a precompile's
// OnStart method), or when a new extension is added.
func (l *globalListenerManager) RegisterNewListener(ctx context.Context, db sql.DB, eng common.Engine, conf EVMEventListenerConfig) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	_, ok := registeredResolutions[conf.Resolve]
	if !ok {
		return fmt.Errorf("resolve function %s not registered", conf.Resolve)
	}

	_, ok = l.listeners[conf.UniqueName]
	if ok {
		return fmt.Errorf("listener with name %s already registered", conf.UniqueName)
	}

	chainInfo, ok := chains.GetChainInfo(conf.Chain)
	if !ok {
		return fmt.Errorf("chain %s not found", conf.Chain)
	}

	err := orderedsync.Synchronizer.RegisterTopic(ctx, db, eng, conf.UniqueName, conf.Resolve)
	if err != nil {
		return err
	}

	doneCh := make(chan struct{})

	linfo := &listenerInfo{
		getLogs:    conf.GetLogs,
		done:       doneCh,
		chain:      chainInfo,
		uniqueName: conf.UniqueName,
	}
	l.listeners[conf.UniqueName] = linfo

	if l.shouldListen {
		// this means it is already listening, so we should start the listener.
		// Otherwise, when the oracle starts, it will listen to all values
		// in the l.listeners map.
		go linfo.listen(l.runningContext, l.runningService, l.runningEventStore, l.runningSyncConf)
	}

	return nil
}

// UnregisterListener unregisters a listener.
// It should be called when an extension gets unused
func (l *globalListenerManager) UnregisterListener(uniqueName string) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	info, ok := l.listeners[uniqueName]
	if !ok {
		return fmt.Errorf("listener with name %s not registered", uniqueName)
	}

	close(info.done)
	delete(l.listeners, uniqueName)

	return nil
}

// listen starts all listeners.
// If it returns an error, the node will stop
func (l *globalListenerManager) listen(ctx context.Context, service *common.Service, eventstore listeners.EventStore, syncConf *syncConfig) error {
	l.mu.Lock()

	if l.shouldListen {
		l.mu.Unlock()
		// protect against making duplicate connections
		return fmt.Errorf("already listening")
	}

	l.shouldListen = true
	l.runningContext = ctx
	l.runningService = service
	l.runningEventStore = eventstore
	l.runningSyncConf = syncConf

	defer func() {
		l.mu.Lock()
		l.shouldListen = false
		// We are simply removing this field in the defer, not modifying the context,
		// so we can ignore the linter warning.
		//nolint:fatcontext
		l.runningContext = nil
		l.runningService = nil
		l.runningEventStore = nil
		l.runningSyncConf = nil
		l.mu.Unlock()
	}()

	for _, info := range l.listeners {
		go info.listen(ctx, service, eventstore, syncConf)
	}

	l.mu.Unlock()
	<-ctx.Done()

	return nil
}

type ResolveFunc func(ctx context.Context, app *common.App, block *common.BlockContext, uniqueName string, logs []ethtypes.Log) error

type listenerInfo struct {
	// done is a channel that is closed when the listener is done
	done chan struct{}
	// chain is the chain that the listener is listening to
	chain chains.ChainInfo
	// uniqueName is the unique name of the listener
	uniqueName string
	// getLogs is a function that queries logs to be synced from the chain
	getLogs GetBlockLogsFunc
}

// listen makes a new client and starts listening for events.
// It does not return any errors because errors returned from listeners
// are fatal, and errors returned from this function are _very_ likely
// due to network errors (e.g. with the target RPC).
func (l *listenerInfo) listen(ctx context.Context, service *common.Service, eventstore listeners.EventStore, syncConf *syncConfig) {
	logger := service.Logger.New(l.uniqueName + "." + string(l.chain.Name))

	chainConf, err := getChainConf(service.LocalConfig.Extensions, l.chain.Name)
	if err != nil {
		logger.Error("failed to get chain config", "err", err)
		return
	}

	ethClient, err := newEthClient(ctx, chainConf.Provider, syncConf.MaxRetries, l.done, logger)
	if err != nil {
		logger.Error("failed to create evm client", "err", err)
		return
	}

	indiv := &individualListener{
		chain:            l.chain,
		syncConf:         syncConf,
		chainConf:        chainConf,
		client:           ethClient,
		orderedSyncTopic: l.uniqueName,
		getLogsFunc:      l.getLogs,
	}

	err = indiv.listen(ctx, eventstore, logger)
	if err != nil {
		logger.Error("error listening", "err", err)
	}
}
