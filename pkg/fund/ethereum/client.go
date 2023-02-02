package ethereum

import (
	"fmt"
	chainClient "kwil/pkg/chain/client"
	chainClientDto "kwil/pkg/chain/client/dto"
	chainClientService "kwil/pkg/chain/client/service"
	"kwil/pkg/fund"
	"kwil/x/contracts/escrow"
	"kwil/x/contracts/token"
)

type Client struct {
	Escrow escrow.EscrowContract
	Token  token.TokenContract

	// TODO: rename this
	ChainClient chainClient.ChainClient

	Config *fund.Config
}

func NewClient(cfg *fund.Config) (*Client, error) {
	chnClient, err := chainClientService.NewChainClientExplicit(&chainClientDto.Config{
		ChainCode:             int64(cfg.ChainCode),
		Endpoint:              cfg.Provider,
		ReconnectionInterval:  cfg.ReConnectionInterval,
		RequiredConfirmations: cfg.RequiredConfirmation,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create chain client: %v", err)
	}

	// escrow
	escrowCtr, err := escrow.New(chnClient, cfg.PrivateKey, cfg.PoolAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to create escrow contract: %v", err)
	}

	// erc20 address
	tokenAddress := escrowCtr.TokenAddress()

	// erc20
	erc20Ctr, err := token.New(chnClient, cfg.PrivateKey, tokenAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to create erc20 contract: %v", err)
	}

	return &Client{
		Escrow:      escrowCtr,
		Token:       erc20Ctr,
		ChainClient: chnClient,
		Config:      cfg,
	}, nil
}
