package service

import (
	"context"

	"github.com/neurowatt/aiwatt-backend/internal/api/dto"
)

// MintServicer defines the mint/redeem business logic interface.
type MintServicer interface {
	Mint(ctx context.Context, req dto.MintRequest) (dto.MintResponse, error)
	Redeem(ctx context.Context, req dto.RedeemRequest) (dto.MintResponse, error)
	GetNAV(ctx context.Context) (dto.NAVResponse, error)
}

// MintService implements MintServicer.
type MintService struct {
	// TODO: inject TxManager, YieldService, Redis NAV cache
}

// NewMintService constructs a MintService.
func NewMintService() *MintService {
	return &MintService{}
}

func (s *MintService) Mint(ctx context.Context, req dto.MintRequest) (dto.MintResponse, error) {
	// TODO: call MintEngine contract via TxManager, track position
	return dto.MintResponse{}, nil
}

func (s *MintService) Redeem(ctx context.Context, req dto.RedeemRequest) (dto.MintResponse, error) {
	// TODO: burn WATT via MintEngine contract, return stablecoin
	return dto.MintResponse{}, nil
}

func (s *MintService) GetNAV(ctx context.Context) (dto.NAVResponse, error) {
	// TODO: read from Redis cache, fallback to YieldService calculation
	return dto.NAVResponse{}, nil
}
