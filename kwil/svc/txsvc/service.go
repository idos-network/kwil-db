package txsvc

import (
	"kwil/kwil/repository"
	"kwil/x/execution/executor"
	"kwil/x/logx"
	"kwil/x/pricing/pricer"
	"kwil/x/proto/txpb"
)

type Service struct {
	txpb.UnimplementedTxServiceServer

	log logx.Logger

	dao repository.Queries

	executor executor.Executor
	pricing  pricer.Pricer
}

func NewService(queries repository.Queries, exec executor.Executor) *Service {
	return &Service{
		log:      logx.New().Named("txsvc"),
		dao:      queries,
		executor: exec,
		pricing:  pricer.NewPricer(),
	}
}
