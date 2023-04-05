package cmd

import (
	"context"
	"fmt"
	"kwil/internal/app/kwild/config"
	"kwil/internal/app/kwild/server"
	"kwil/internal/controller/grpc/healthsvc/v0"
	"kwil/internal/controller/grpc/txsvc/v1"
	"kwil/internal/pkg/gateway/middleware/cors"
	"kwil/internal/pkg/healthcheck"
	simple_checker "kwil/internal/pkg/healthcheck/simple-checker"

	"kwil/pkg/balances"
	chainsyncer "kwil/pkg/balances/chain-syncer"
	chainClient "kwil/pkg/chain/client"
	ccService "kwil/pkg/chain/client/service" // shorthand for chain client service
	chainTypes "kwil/pkg/chain/types"
	kwilCrypto "kwil/pkg/crypto"
	"kwil/pkg/log"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var RootCmd = &cobra.Command{
	Use:   "kwild",
	Short: "kwil grpc server",
	Long:  "",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()
		cfg, err := config.LoadConfig()
		if err != nil {
			return err
		}

		logger := log.New(cfg.Log)

		ctx, cancel := context.WithCancel(ctx)
		defer cancel()

		chainclient, err := buildChainClient(cfg, logger)
		if err != nil {
			return fmt.Errorf("failed to build chain client: %w", err)
		}

		accountStore, err := buildAccountRepository(logger, cfg)
		if err != nil {
			return fmt.Errorf("failed to build account repository: %w", err)
		}

		chainSyncer, err := buildChainSyncer(cfg, chainclient, accountStore, logger)
		if err != nil {
			return fmt.Errorf("failed to build chain syncer: %w", err)
		}

		txSvc, err := buildTxSvc(cfg, accountStore, logger)
		if err != nil {
			return fmt.Errorf("failed to build tx service: %w", err)
		}

		// kwil gateway
		gw := server.NewGWServer(runtime.NewServeMux(), cfg, logger)

		if err := gw.SetupGrpcSvc(ctx); err != nil {
			return err
		}
		if err := gw.SetupHTTPSvc(ctx); err != nil {
			return err
		}

		gw.AddMiddlewares(
			// from innermost middleware
			//auth.MAuth(keyManager, logger),
			cors.MCors([]string{}),
		)

		server := &server.Server{
			Cfg:         cfg,
			Log:         logger,
			ChainSyncer: chainSyncer,
			TxSvc:       txSvc,
			HealthSvc:   buildHealthSvc(logger),
		}

		go func() {
			if err := gw.Serve(); err != nil {
				logger.Error("failed to serve kwil gateway", zap.Error(err))
			}
		}()

		return server.Start(ctx)
	}}

func init() {
	/*
		defaultConfigPath := filepath.Join("$HOME", config.DefaultConfigDir,
			fmt.Sprintf("%s.%s", config.DefaultConfigName, config.DefaultConfigType))
		RootCmd.PersistentFlags().StringVar(&config.ConfigFile, "config", "", fmt.Sprintf("config file to use (default: '%s')", defaultConfigPath))
	*/

	config.BindGlobalFlags(RootCmd.PersistentFlags())
	config.BindGlobalEnv(RootCmd.PersistentFlags())
}

func buildChainClient(cfg *config.KwildConfig, logger log.Logger) (chainClient.ChainClient, error) {
	return ccService.NewChainClient(cfg.Deposits.ClientChainRPCURL,
		ccService.WithChainCode(chainTypes.ChainCode(cfg.Deposits.ChainCode)),
		ccService.WithLogger(logger),
		ccService.WithReconnectInterval(int64(cfg.Deposits.ReconnectionInterval)),
		ccService.WithRequiredConfirmations(int64(cfg.Deposits.BlockConfirmations)),
	)
}

func buildAccountRepository(logger log.Logger, cfg *config.KwildConfig) (*balances.AccountStore, error) {
	return balances.NewAccountStore(
		balances.WithLogger(logger),
		balances.WithPath(cfg.SqliteFilePath),
	)
}

func buildChainSyncer(cfg *config.KwildConfig, cc chainClient.ChainClient, as *balances.AccountStore, logger log.Logger) (*chainsyncer.ChainSyncer, error) {
	walletAddress := kwilCrypto.AddressFromPrivateKey(cfg.PrivateKey)

	return chainsyncer.Builder().
		WithLogger(logger).
		WritesTo(as).
		ListensTo(cfg.Deposits.PoolAddress).
		WithChainClient(cc).
		WithReceiverAddress(walletAddress).
		Build()
}

func buildTxSvc(cfg *config.KwildConfig, as *balances.AccountStore, logger log.Logger) (*txsvc.Service, error) {
	return txsvc.NewService(cfg,
		txsvc.WithLogger(logger),
		txsvc.WithAccountStore(as),
	)
}

func buildHealthSvc(logger log.Logger) *healthsvc.Server {
	// health service
	registrar := healthcheck.NewRegistrar(logger)
	registrar.RegisterAsyncCheck(10*time.Second, 3*time.Second, healthcheck.Check{
		Name: "dummy",
		Check: func(ctx context.Context) error {
			// error make this check fail, nil will make it succeed
			return nil
		},
	})
	ck := registrar.BuildChecker(simple_checker.New(logger))
	return healthsvc.NewServer(ck)
}

// from v0, removed 04/03/23
/*
	ctx := cmd.Context()
	cfg, err := config.LoadConfig()
	if err != nil {
		return err
	}

	// build log
	//log, err := log.NewLogger(cfg.log)
	logger := log.New(cfg.Log)

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	client, err := sqlclient.Open(cfg.DB.DbUrl(), 60*time.Second)
	if err != nil {
		return fmt.Errorf("failed to open sql client: %w", err)
	}

	//&cfg.Fund.Chain, logger
	chainClient, err := service.NewChainClient(cfg.Fund.Chain.RpcUrl,
		service.WithChainCode(chainTypes.ChainCode(cfg.Fund.Chain.ChainCode)),
		service.WithLogger(logger),
		service.WithReconnectInterval(cfg.Fund.Chain.ReconnectInterval),
		service.WithRequiredConfirmations(cfg.Fund.Chain.BlockConfirmation),
	)
	if err != nil {
		return fmt.Errorf("failed to build chain client: %w", err)
	}

	// build repository prepared statement
	queries, err := repository.Prepare(ctx, client)
	if err != nil {
		return fmt.Errorf("failed to prepare queries: %w", err)
	}

	dps, err := deposits.NewDepositer(cfg.Fund.PoolAddress, client, queries, chainClient, cfg.Fund.Wallet, logger)
	if err != nil {
		return fmt.Errorf("failed to build deposits: %w", err)
	}

	hasuraManager := hasura.NewClient(cfg.Graphql.Addr, logger)
	go hasura.Initialize(cfg.Graphql.Addr, logger)

	// build executor
	exec, err := executor.NewExecutor(ctx, client, queries, hasuraManager, logger)
	if err != nil {
		return fmt.Errorf("failed to build executor: %w", err)
	}

	// build config service
	accSvc := accountsvc.NewService(queries, logger)

	// pricing service
	prcSvc := pricingsvc.NewService(exec)

	// tx service
	txService := txsvc.NewService(queries, exec, logger)

	// health service
	registrar := healthcheck.NewRegistrar(logger)
	registrar.RegisterAsyncCheck(10*time.Second, 3*time.Second, healthcheck.Check{
		Name: "dummy",
		Check: func(ctx context.Context) error {
			// error make this check fail, nil will make it succeed
			return nil
		},
	})
	ck := registrar.BuildChecker(simple_checker.New(logger))
	healthService := healthsvc.NewServer(ck)

	// configuration service
	cfgService := configsvc.NewService(cfg, logger)
	// build server
	svr := server.New(cfg.Server, txService, accSvc, cfgService, healthService, prcSvc, dps, logger)
	return svr.Start(ctx)
*/