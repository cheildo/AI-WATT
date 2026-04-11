package service

import "context"

// TreasuryServicer implements Engine 3 — idle capital auto-deployment.
type TreasuryServicer interface {
	SweepIdleCapital(ctx context.Context) error
	ReportFeeAccrual(ctx context.Context) (float64, error)
}

// TreasuryService implements TreasuryServicer.
type TreasuryService struct {
	// TODO: inject TxManager, YieldService
}

func NewTreasuryService() *TreasuryService { return &TreasuryService{} }

func (s *TreasuryService) SweepIdleCapital(ctx context.Context) error {
	// TODO: detect idle capital in MintEngine, deploy to T-bill wrapper
	return nil
}

func (s *TreasuryService) ReportFeeAccrual(ctx context.Context) (float64, error) {
	// TODO: sum protocol fees + T-bill yield, report to YieldService
	return 0, nil
}
