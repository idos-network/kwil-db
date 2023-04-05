package kwild

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"kwil/internal/pkg/graphql/query"
	"kwil/pkg/client"
	"kwil/pkg/engine/models"
	grpc "kwil/pkg/grpc/client/v1"
	"kwil/pkg/log"
	kTx "kwil/pkg/tx"
	"math/big"
	"strings"

	ec "github.com/ethereum/go-ethereum/crypto"
	"go.uber.org/zap"
)

// KwildDriver is a grpc driver for  integration tests
type KwildDriver struct {
	clt         *client.Client
	pk          *ecdsa.PrivateKey
	gatewayAddr string // to ignore the gatewayAddr returned by the config.service

	logger log.Logger
}

func NewKwildDriver(clt *client.Client, pk *ecdsa.PrivateKey, gatewayAddr string, logger log.Logger) *KwildDriver {
	return &KwildDriver{
		clt:         clt,
		pk:          pk,
		gatewayAddr: gatewayAddr,
		logger:      logger,
	}
}

func (d *KwildDriver) GetUserAddress() string {
	return ec.PubkeyToAddress(d.pk.PublicKey).Hex()
}

func (d *KwildDriver) GetServiceConfig(ctx context.Context) (*grpc.SvcConfig, error) {
	return d.clt.GetConfig(ctx)
}

func (d *KwildDriver) DepositFund(ctx context.Context, amount *big.Int) error {
	_, err := d.clt.Deposit(ctx, amount)
	if err != nil {
		return fmt.Errorf("failed to send deposit transaction: %w", err)
	}

	d.logger.Debug("deposit fund", zap.String("from", d.GetUserAddress()),
		zap.String("to", d.clt.ProviderAddress), zap.String("amount", amount.String()))
	return nil
}

func (d *KwildDriver) GetDepositBalance(ctx context.Context) (*big.Int, error) {
	bal, err := d.clt.GetDepositBalance(ctx)
	if err != nil {
		return nil, err
	}
	return bal, nil
}

func (d *KwildDriver) ApproveToken(ctx context.Context, amount *big.Int) error {
	_, err := d.clt.ApproveDeposit(ctx, amount)
	if err != nil {
		return err
	}

	d.logger.Debug("approve token", zap.String("from", ec.PubkeyToAddress(d.pk.PublicKey).Hex()),
		zap.String("spender", d.clt.PoolAddress), zap.String("amount", amount.String()))
	return nil
}

func (d *KwildDriver) GetAllowance(ctx context.Context) (*big.Int, error) {
	amount, err := d.clt.GetApprovedAmount(ctx)
	if err != nil {
		return nil, err
	}

	return amount, nil
}

func (d *KwildDriver) DeployDatabase(ctx context.Context, db *models.Dataset) error {
	_, err := d.clt.DeployDatabase(ctx, db)
	if err != nil {
		return fmt.Errorf("error deploying database: %w", err)
	}

	d.logger.Debug("deploy database", zap.String("name", db.Name), zap.String("owner", db.Owner))
	return nil
}

func (d *KwildDriver) DatabaseShouldExists(ctx context.Context, owner string, dbName string) error {
	dbid := models.GenerateSchemaId(owner, dbName)

	schema, err := d.clt.GetSchema(ctx, dbid)
	if err != nil {
		return fmt.Errorf("failed to get database schema: %w", err)
	}

	if strings.EqualFold(schema.Owner, owner) && strings.EqualFold(schema.Name, dbName) {
		return nil
	}
	return fmt.Errorf("database does not exist")
}

func (d *KwildDriver) ExecuteAction(ctx context.Context, dbid string, queryName string, queryInputs []map[string]any) (*kTx.Receipt, error) {
	rec, err := d.clt.ExecuteAction(ctx, dbid, queryName, queryInputs)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %w", err)
	}

	d.logger.Debug("execute query", zap.String("database", dbid), zap.String("query", queryName))
	return rec, nil
}

func (d *KwildDriver) DropDatabase(ctx context.Context, dbName string) error {
	_, err := d.clt.DropDatabase(ctx, dbName)
	if err != nil {
		return fmt.Errorf("error dropping database: %w", err)
	}
	d.logger.Debug("drop database", zap.String("name", dbName), zap.String("owner", d.GetUserAddress()))
	return nil
}

func (d *KwildDriver) QueryDatabase(ctx context.Context, queryStr string) ([]byte, error) {
	url := fmt.Sprintf("http://%s/graphql", d.gatewayAddr)
	return query.Query(ctx, url, queryStr)
}