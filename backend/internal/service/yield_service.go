package service

import "context"

// YieldServicer calculates NAV and distributes yield to sWATT holders.
type YieldServicer interface {
	ComputeNAV(ctx context.Context) (float64, error)
	DistributeYield(ctx context.Context, loanID string, amount float64) error
}

// YieldService implements YieldServicer.
type YieldService struct {
	// TODO: inject LoanRepo, Redis cache
}

func NewYieldService() *YieldService { return &YieldService{} }

func (s *YieldService) ComputeNAV(ctx context.Context) (float64, error) {
	// TODO: totalAssets / totalSupply from on-chain + DB
	return 1.0, nil
}

func (s *YieldService) DistributeYield(ctx context.Context, loanID string, amount float64) error {
	// TODO: update sWATT vault accounting, cache new NAV in Redis
	return nil
}
