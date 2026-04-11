package service

import "context"

// WEVServicer manages the sWATT redemption queue.
type WEVServicer interface {
	RequestRedeem(ctx context.Context, userID string, amount float64) error
	ProcessQueue(ctx context.Context) error
}

// WEVService implements WEVServicer.
type WEVService struct {
	// TODO: inject WEVRepo, TxManager
}

func NewWEVService() *WEVService { return &WEVService{} }

func (s *WEVService) RequestRedeem(ctx context.Context, userID string, amount float64) error {
	// TODO: insert into wev_queue table, call WEVQueue contract
	return nil
}

func (s *WEVService) ProcessQueue(ctx context.Context) error {
	// TODO: daily batch processing of redemption queue
	return nil
}
