// Package main runs a node stress test tool with a few programs designed to
// impose a high load and test edge cases in transaction handling and dataset
// engine operations such as dataset deployment and action execution.

package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

var (
	host            string
	gatewayProvider bool
	quiet           bool

	numClients int

	// namespace is the existing namespace to use (no need to be DB owner),
	// otherwise deploy a new one (requires DB owner key)
	namespace string

	chainId  string
	authCall bool

	runTime time.Duration

	badgerInterval time.Duration
	viewInterval   time.Duration

	deployDropInterval time.Duration
	fastDropRate       int
	noDrop             bool

	noErrActs bool

	maxPosters   int
	postInterval time.Duration
	contentLen   int
	variableLen  bool

	txPollInterval time.Duration

	concurrentBroadcast bool
	nonceChaos          int
	rpcTiming           bool

	wg sync.WaitGroup
)

func main() {
	var key string

	flag.StringVar(&host, "host", "http://127.0.0.1:8484", "provider's http url, schema is required")
	flag.BoolVar(&gatewayProvider, "gw", false, "gateway provider instead of vanilla provider, "+
		"need to make sure host is same as gateway's domain")
	flag.StringVar(&key, "key", "", "existing key to use instead of generating a new one, only applies to the first client")
	flag.IntVar(&numClients, "cl", 1, "number of simulated clients (with unique keys)")
	flag.StringVar(&namespace, "ns", "", "existing namespace to use instead of deploying a new one, which would require DB owner key")
	flag.BoolVar(&quiet, "q", false, "only print errors")

	flag.StringVar(&chainId, "chain", "", "chain ID to require (default is any)")
	flag.BoolVar(&authCall, "authcall", false, "sign our call requests expecting the server to authenticate (private mode)")

	flag.DurationVar(&runTime, "run", 30*time.Minute, "terminate after running this long")

	flag.DurationVar(&badgerInterval, "bi", -1, "badger kwild with read-only metadata requests at this interval")

	flag.DurationVar(&deployDropInterval, "ddi", -1, "deploy/drop datasets at this interval (but after drop tx confirms)")
	flag.IntVar(&fastDropRate, "ddn", 0, "immediately drop new dbs at a rate of 1/ddn (disable with <1)")
	flag.BoolVar(&noDrop, "nodrop", false, "don't drop the datasets deployed in the deploy/drop program")

	flag.BoolVar(&noErrActs, "ne", false, "don't make intentionally failed txns")

	flag.IntVar(&maxPosters, "ec", 4, "max concurrent unconfirmed action executions (to get multiple tx in a block)")
	flag.DurationVar(&postInterval, "ei", 10*time.Millisecond,
		"initiate non-view action execution at this interval (limited by max concurrency setting)")
	flag.DurationVar(&viewInterval, "vi", -1, "make view action call at this interval")
	flag.IntVar(&contentLen, "el", 50_000, "content length in an executed post action")
	flag.BoolVar(&variableLen, "vl", false, "pseudorandom variable content lengths, on (0,el]")

	flag.BoolVar(&concurrentBroadcast, "cb", false, "concurrent broadcast (do not wait for broadcast result before releasing nonce lock, will cause nonce errors due to goroutine race)")
	flag.IntVar(&nonceChaos, "nc", 0, "nonce chaos rate (apply nonce jitter every 1/nc times)")
	flag.BoolVar(&rpcTiming, "v", false, "print RPC durations")

	flag.DurationVar(&txPollInterval, "pollint", 400*time.Millisecond, "polling interval when waiting for tx confirmation")

	flag.Parse()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	complete := errors.New("reached end time")
	ctx, cancel := context.WithTimeoutCause(context.Background(), runTime, complete)

	go func() {
		<-signalChan
		cancel()
	}()

	errChan := make(chan error, 1)
	dbReady := make(chan struct{})

	var wg sync.WaitGroup
	for i := range numClients {
		wg.Add(1)
		go func(key string) {
			defer wg.Done()
			errChan <- hammer(ctx, key, strconv.Itoa(i), dbReady)
		}(key)

		if key != "" {
			key = ""
		}

		if i == 0 {
			select {
			case <-dbReady:
				dbReady = nil // ready for additional clients
			case err := <-errChan:
				fmt.Println(err) // setup failed
				os.Exit(1)
			}
		}
	}

	go func() {
		wg.Wait()
		close(errChan)
	}()

	var exitCode int
	for err := range errChan {
		if err != nil && !errors.Is(err, complete) {
			fmt.Println(err)
			exitCode = 1
		}
	}

	cancel()

	os.Exit(exitCode)
}
