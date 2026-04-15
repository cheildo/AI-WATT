package service

import (
	"context"
	"time"

	"github.com/neurowatt/aiwatt-backend/internal/api/dto"
)

// YieldServicer calculates NAV and vault statistics.
type YieldServicer interface {
	GetVaultStats(ctx context.Context) (dto.VaultStatsResponse, error)
}

// YieldService implements YieldServicer.
// Phase 8 will wire in BlockchainClient reads from sWattUSD for real on-chain NAV.
type YieldService struct{}

func NewYieldService() *YieldService { return &YieldService{} }

func (s *YieldService) GetVaultStats(ctx context.Context) (dto.VaultStatsResponse, error) {
	// Placeholder — Phase 8 replaces with on-chain sWattUSD.navPerShare() + totalAssets().
	return dto.VaultStatsResponse{
		NAVPerShare:   1.0,
		TotalAssets:   0,
		TotalSupply:   0,
		DeployedPct:   0,
		TBillReserve:  0,
		APR7D:         0,
		APR30D:        0,
		LastUpdatedAt: time.Now().UTC().Format(time.RFC3339),
	}, nil
}
