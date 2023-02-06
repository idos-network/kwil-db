package kwil_client

import (
	"kwil/pkg/fund"
	"kwil/pkg/grpc/client"
	"kwil/pkg/logger"
)

type Config struct {
	// Node config
	// @yaiba TODO: a better name, maybe Peer?
	Node client.Config `mapstructure:"node"`
	// Fund config
	// @yaiba TODO: a better name, maybe SettlementChain?
	Fund fund.Config `mapstructure:"fund"`

	Log logger.Config `mapstructure:"log"`
}
