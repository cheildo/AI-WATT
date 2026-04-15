package service

import (
	"context"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/google/uuid"
	"go.uber.org/zap"

	"github.com/neurowatt/aiwatt-backend/internal/api/dto"
	"github.com/neurowatt/aiwatt-backend/internal/blockchain"
	"github.com/neurowatt/aiwatt-backend/internal/models"
	"github.com/neurowatt/aiwatt-backend/internal/repository"
)

// AssetServicer defines the asset business logic interface.
type AssetServicer interface {
	Register(ctx context.Context, req dto.RegisterAssetRequest) (dto.AssetResponse, error)
	GetByID(ctx context.Context, id string) (dto.AssetResponse, error)
	UpdateLTV(ctx context.Context, id string, newLTV float64) (dto.AssetResponse, error)
	List(ctx context.Context, q dto.ListAssetsQuery) (dto.ListAssetsResponse, error)
	GetHealth(ctx context.Context, assetID string) (dto.HealthScoreResponse, error)
}

// AssetService implements AssetServicer.
type AssetService struct {
	assetRepo repository.AssetRepo
	txManager *blockchain.TxManager
	logger    *zap.Logger
}

func NewAssetService(assetRepo repository.AssetRepo, txManager *blockchain.TxManager, logger *zap.Logger) *AssetService {
	return &AssetService{assetRepo: assetRepo, txManager: txManager, logger: logger}
}

func (s *AssetService) Register(ctx context.Context, req dto.RegisterAssetRequest) (dto.AssetResponse, error) {
	asset := &models.Asset{
		ID:        uuid.NewString(),
		AssetType: req.AssetType,
		OwnerID:   req.OwnerID,
		Status:    models.AssetStatusPending,
	}
	if err := s.assetRepo.Create(ctx, asset); err != nil {
		return dto.AssetResponse{}, fmt.Errorf("asset_service.Register: %w", err)
	}

	// Register on-chain and mint OC-NFT asynchronously so the HTTP response is fast.
	go s.registerOnChain(asset, req)

	return toAssetDTO(asset), nil
}

// registerOnChain calls AssetRegistry.registerAsset and OCNFT.mintOCNFT on XDC.
// Runs in a goroutine — failures are logged but do not block the API response.
func (s *AssetService) registerOnChain(asset *models.Asset, req dto.RegisterAssetRequest) {
	ctx := context.Background()
	assetID32 := uuidToBytes32(asset.ID)
	borrower := common.HexToAddress(req.BorrowerWallet)
	ltvBPS := uint16(req.InitialLTV * 10000)
	assetTypeUint := assetTypeToUint8(req.AssetType)

	if _, err := s.txManager.RegisterAssetOnChain(ctx, assetID32, assetTypeUint, borrower, ltvBPS); err != nil {
		s.logger.Error("asset_service.registerOnChain: RegisterAssetOnChain failed",
			zap.String("asset_id", asset.ID),
			zap.Error(err),
		)
		return
	}

	// metadataURI will be replaced by a real IPFS CID in Phase 9 (Pinata upload).
	metadataURI := fmt.Sprintf("ipfs://pending/%s", asset.ID)
	if _, err := s.txManager.MintOCNFT(ctx, borrower, assetID32, metadataURI); err != nil {
		s.logger.Error("asset_service.registerOnChain: MintOCNFT failed",
			zap.String("asset_id", asset.ID),
			zap.Error(err),
		)
	}
}

func (s *AssetService) GetByID(ctx context.Context, id string) (dto.AssetResponse, error) {
	asset, err := s.assetRepo.GetByID(ctx, id)
	if err != nil {
		return dto.AssetResponse{}, fmt.Errorf("asset_service.GetByID: %w", err)
	}
	return toAssetDTO(asset), nil
}

func (s *AssetService) UpdateLTV(ctx context.Context, id string, newLTV float64) (dto.AssetResponse, error) {
	asset, err := s.assetRepo.GetByID(ctx, id)
	if err != nil {
		return dto.AssetResponse{}, fmt.Errorf("asset_service.UpdateLTV: %w", err)
	}
	asset.CurrentLTV = newLTV
	if err := s.assetRepo.Update(ctx, asset); err != nil {
		return dto.AssetResponse{}, fmt.Errorf("asset_service.UpdateLTV: save: %w", err)
	}

	// Update on-chain asynchronously.
	go func() {
		ltvBPS := uint16(newLTV * 10000)
		if _, err := s.txManager.UpdateLTVOnChain(context.Background(), uuidToBytes32(id), ltvBPS); err != nil {
			s.logger.Error("asset_service.UpdateLTV: UpdateLTVOnChain failed",
				zap.String("asset_id", id),
				zap.Error(err),
			)
		}
	}()

	return toAssetDTO(asset), nil
}

func (s *AssetService) List(ctx context.Context, q dto.ListAssetsQuery) (dto.ListAssetsResponse, error) {
	offset := (q.Page - 1) * q.PageSize
	assets, total, err := s.assetRepo.List(ctx, q.OwnerID, q.Status, offset, q.PageSize)
	if err != nil {
		return dto.ListAssetsResponse{}, fmt.Errorf("asset_service.List: %w", err)
	}
	resp := make([]dto.AssetResponse, len(assets))
	for i, a := range assets {
		resp[i] = toAssetDTO(a)
	}
	return dto.ListAssetsResponse{Assets: resp, Total: total, Page: q.Page}, nil
}

func (s *AssetService) GetHealth(ctx context.Context, assetID string) (dto.HealthScoreResponse, error) {
	asset, err := s.assetRepo.GetByID(ctx, assetID)
	if err != nil {
		return dto.HealthScoreResponse{}, fmt.Errorf("asset_service.GetHealth: %w", err)
	}
	return dto.HealthScoreResponse{
		AssetID:     asset.ID,
		HealthScore: asset.HealthScore,
		Status:      asset.Status,
		ComputedAt:  asset.UpdatedAt.Format(time.RFC3339),
	}, nil
}

// ── Helpers ───────────────────────────────────────────────────────────────

func toAssetDTO(a *models.Asset) dto.AssetResponse {
	return dto.AssetResponse{
		ID:           a.ID,
		AssetType:    a.AssetType,
		LoanID:       a.LoanID,
		OwnerID:      a.OwnerID,
		HealthScore:  a.HealthScore,
		CurrentLTV:   a.CurrentLTV,
		Status:       a.Status,
		OCNFTTokenID: a.OCNFTTokenID,
		MetadataURI:  a.MetadataURI,
		CreatedAt:    a.CreatedAt.Format(time.RFC3339),
	}
}

// uuidToBytes32 derives a deterministic bytes32 on-chain ID from a UUID string
// by keccak256-hashing its bytes.
func uuidToBytes32(id string) [32]byte {
	hash := crypto.Keccak256([]byte(id))
	var b [32]byte
	copy(b[:], hash)
	return b
}

// assetTypeToUint8 maps an asset type string to the on-chain enum value.
func assetTypeToUint8(t string) uint8 {
	switch t {
	case models.AssetTypeGPUCluster:
		return 0
	case models.AssetTypeRobotics:
		return 1
	case models.AssetTypeEnergy:
		return 2
	default:
		return 0
	}
}
