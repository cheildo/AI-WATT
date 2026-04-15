package service

import (
	"context"
	"time"

	"github.com/neurowatt/aiwatt-backend/internal/api/dto"
)

// MintServicer defines the mint/redeem business logic interface.
type MintServicer interface {
	Mint(ctx context.Context, req dto.MintRequest) (dto.MintResponse, error)
	Redeem(ctx context.Context, req dto.RedeemRequest) (dto.MintResponse, error)
	GetNAV(ctx context.Context) (dto.NAVResponse, error)
}

// MintService implements MintServicer.
// Phase 8 will wire in TxManager calls to the MintEngine contract.
type MintService struct{}

func NewMintService() *MintService { return &MintService{} }

func (s *MintService) Mint(ctx context.Context, req dto.MintRequest) (dto.MintResponse, error) {
	// Phase 8: call MintEngine.depositAndMint via TxManager; return real tx hash + amounts.
	return dto.MintResponse{AmountMinted: req.Amount, Fee: req.Amount * 0.001}, nil
}

func (s *MintService) Redeem(ctx context.Context, req dto.RedeemRequest) (dto.MintResponse, error) {
	// Phase 8: call MintEngine.redeemWATT via TxManager.
	return dto.MintResponse{AmountRedeemed: req.Amount}, nil
}

func (s *MintService) GetNAV(ctx context.Context) (dto.NAVResponse, error) {
	// Phase 8: read from Redis cache, fallback to BlockchainClient.navPerShare().
	return dto.NAVResponse{
		NAVPerShare:   1.0,
		LastUpdatedAt: time.Now().UTC().Format(time.RFC3339),
	}, nil
}
