package service

import (
	"context"

	"github.com/neurowatt/aiwatt-backend/internal/api/dto"
)

// AssetServicer defines the asset business logic interface.
type AssetServicer interface {
	Register(ctx context.Context, req dto.RegisterAssetRequest) (dto.AssetResponse, error)
	GetByID(ctx context.Context, id string) (dto.AssetResponse, error)
	UpdateLTV(ctx context.Context, id string, newLTV float64) (dto.AssetResponse, error)
	List(ctx context.Context, q dto.ListAssetsQuery) (dto.ListAssetsResponse, error)
}

// AssetService implements AssetServicer.
type AssetService struct {
	// TODO: inject AssetRepo, TxManager (for OCNFT mint + AssetRegistry call), IPFS client
}

// NewAssetService constructs an AssetService.
func NewAssetService() *AssetService {
	return &AssetService{}
}

func (s *AssetService) Register(ctx context.Context, req dto.RegisterAssetRequest) (dto.AssetResponse, error) {
	// TODO: create asset, upload metadata to IPFS, register on AssetRegistry contract
	return dto.AssetResponse{}, nil
}

func (s *AssetService) GetByID(ctx context.Context, id string) (dto.AssetResponse, error) {
	// TODO: fetch from AssetRepo, map model → dto
	return dto.AssetResponse{}, nil
}

func (s *AssetService) UpdateLTV(ctx context.Context, id string, newLTV float64) (dto.AssetResponse, error) {
	// TODO: update LTV in DB and on-chain via AssetRegistry contract
	return dto.AssetResponse{}, nil
}

func (s *AssetService) List(ctx context.Context, q dto.ListAssetsQuery) (dto.ListAssetsResponse, error) {
	// TODO: paginated fetch from AssetRepo
	return dto.ListAssetsResponse{}, nil
}
