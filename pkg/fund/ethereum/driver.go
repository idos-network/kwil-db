package ethereum

import (
	"context"
	"fmt"
	"kwil/pkg/fund"
	"kwil/pkg/logger"
	"math/big"
	"sync"
)

// Driver is a driver for the chain client for integration tests
type Driver struct {
	RpcUrl string

	logger     logger.Logger
	connOnce   sync.Once
	Fund       fund.IFund
	fundConfig *fund.Config
}

func New(rpcUrl string, logger logger.Logger) *Driver {
	return &Driver{
		logger: logger,
		RpcUrl: rpcUrl,
	}
}

func (d *Driver) DepositFund(ctx context.Context, to string, amount *big.Int) error {
	client, err := d.GetClient()
	if err != nil {
		return fmt.Errorf("failed to create client: %w", err)
	}

	_, err = client.DepositFund(ctx, to, amount)

	return err
}

func (d *Driver) GetDepositBalance(ctx context.Context, from string, to string) (*big.Int, error) {
	client, err := d.GetClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %w", err)
	}

	return client.GetDepositBalance(ctx, from, to)
}

func (d *Driver) ApproveToken(ctx context.Context, spender string, amount *big.Int) error {
	client, err := d.GetClient()
	if err != nil {
		return fmt.Errorf("failed to create client: %w", err)
	}

	_, err = client.ApproveToken(ctx, spender, amount)

	return err
}

func (d *Driver) GetAllowance(ctx context.Context, from string, spender string) (*big.Int, error) {
	client, err := d.GetClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %w", err)
	}

	return client.GetAllowance(ctx, from, spender)
}

func (d *Driver) GetFundConfig() *fund.Config {
	return d.fundConfig
}

func (d *Driver) SetFundConfig(cfg *fund.Config) {
	d.fundConfig = cfg
}

func (d *Driver) GetClient() (fund.IFund, error) {
	var err error
	d.connOnce.Do(func() {
		d.Fund, err = NewClient(d.fundConfig, d.logger)
	})

	return d.Fund, err
}
