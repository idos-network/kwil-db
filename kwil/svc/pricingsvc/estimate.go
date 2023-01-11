package pricingsvc

import (
	"context"
	"fmt"
	"kwil/x/proto/commonpb"
	"kwil/x/proto/pricingpb"
	"kwil/x/types/transactions"
	"kwil/x/utils/serialize"
)

func (s *Service) EstimateCost(ctx context.Context, req *pricingpb.EstimateRequest) (*pricingpb.EstimateResponse, error) {
	tx, err := serialize.Convert[commonpb.Tx, transactions.Transaction](req.Tx)
	if err != nil {
		return nil, fmt.Errorf("failed to convert transaction: %w", err)
	}

	price, err := s.pricer.EstimatePrice(ctx, tx)
	if err != nil {
		return nil, fmt.Errorf("failed to estimate price: %w", err)
	}

	return &pricingpb.EstimateResponse{
		Cost: price,
	}, nil
}
