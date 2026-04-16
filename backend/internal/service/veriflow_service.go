package service

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"go.uber.org/zap"

	"github.com/neurowatt/aiwatt-backend/internal/api/dto"
	"github.com/neurowatt/aiwatt-backend/internal/blockchain"
	"github.com/neurowatt/aiwatt-backend/internal/models"
	"github.com/neurowatt/aiwatt-backend/internal/repository"
	"github.com/neurowatt/aiwatt-backend/internal/veriflow"
	"github.com/neurowatt/aiwatt-backend/pkg/crypto"
)

// onChainFlagged is the AssetStatus enum value for FLAGGED in AssetRegistry.sol (PENDING=0, ACTIVE=1, FLAGGED=2).
const onChainFlagged uint8 = 2

// VeriflowServicer defines the Veriflow telemetry + scoring interface.
type VeriflowServicer interface {
	IngestTelemetry(ctx context.Context, payload dto.TelemetryPayload) error
	GetHealthScore(ctx context.Context, assetID string) (dto.HealthScoreResponse, error)
	GetAttestation(ctx context.Context, assetID string) (dto.AttestationResponse, error)
}

// VeriflowService implements VeriflowServicer.
type VeriflowService struct {
	telemetryRepo   repository.TelemetryRepo
	attestationRepo repository.AttestationRepo
	assetRepo       repository.AssetRepo
	scorer          *veriflow.Scorer
	txManager       *blockchain.TxManager
	notifySvc       NotifyServicer
	logger          *zap.Logger
}

func NewVeriflowService(
	telemetryRepo repository.TelemetryRepo,
	attestationRepo repository.AttestationRepo,
	assetRepo repository.AssetRepo,
	scorer *veriflow.Scorer,
	txManager *blockchain.TxManager,
	notifySvc NotifyServicer,
	logger *zap.Logger,
) *VeriflowService {
	return &VeriflowService{
		telemetryRepo:   telemetryRepo,
		attestationRepo: attestationRepo,
		assetRepo:       assetRepo,
		scorer:          scorer,
		txManager:       txManager,
		notifySvc:       notifySvc,
		logger:          logger,
	}
}

// IngestTelemetry verifies the HMAC signature, validates asset status, stores the
// telemetry row, and fires async health scoring. The HMAC secret is fetched from
// assets.hmac_secret provisioned at asset onboarding.
func (s *VeriflowService) IngestTelemetry(ctx context.Context, payload dto.TelemetryPayload) error {
	// 1. Look up asset — ensures it exists and retrieves the HMAC secret.
	asset, err := s.assetRepo.GetByID(ctx, payload.AssetID)
	if err != nil {
		return fmt.Errorf("veriflow_service.IngestTelemetry: asset not found: %w", err)
	}

	// 2. Validate asset is ACTIVE.
	if asset.Status != models.AssetStatusActive {
		return fmt.Errorf("veriflow_service.IngestTelemetry: asset %s is not ACTIVE (status=%s)", payload.AssetID, asset.Status)
	}

	// 3. Verify HMAC — sign JSON of all fields except hmac_signature.
	payloadBytes, err := json.Marshal(dto.TelemetryPayload{
		AssetID:         payload.AssetID,
		Timestamp:       payload.Timestamp,
		GPUUtilization:  payload.GPUUtilization,
		GPUTemperature:  payload.GPUTemperature,
		GPUMemoryUsedMB: payload.GPUMemoryUsedMB,
		GPUErrorRate:    payload.GPUErrorRate,
		ECCErrors:       payload.ECCErrors,
		PowerDrawWatts:  payload.PowerDrawWatts,
		FanSpeedRPM:     payload.FanSpeedRPM,
		UptimePct:       payload.UptimePct,
	})
	if err != nil {
		return fmt.Errorf("veriflow_service.IngestTelemetry: marshal: %w", err)
	}
	if !crypto.HMACVerify([]byte(asset.HMACSecret), payloadBytes, payload.HMACSignature) {
		return fmt.Errorf("veriflow_service.IngestTelemetry: HMAC mismatch for asset %s", payload.AssetID)
	}

	// 4. Persist telemetry row.
	row := &models.Telemetry{
		AssetID:         payload.AssetID,
		GPUUtilization:  payload.GPUUtilization,
		GPUTemperature:  payload.GPUTemperature,
		GPUMemoryUsedMB: payload.GPUMemoryUsedMB,
		GPUErrorRate:    payload.GPUErrorRate,
		ECCErrors:       payload.ECCErrors,
		PowerDrawWatts:  payload.PowerDrawWatts,
		FanSpeedRPM:     payload.FanSpeedRPM,
		UptimePct:       payload.UptimePct,
		HMACSignature:   payload.HMACSignature,
		RecordedAt:      time.Unix(payload.Timestamp, 0).UTC(),
	}
	if err := s.telemetryRepo.Create(ctx, row); err != nil {
		return fmt.Errorf("veriflow_service.IngestTelemetry: insert: %w", err)
	}

	// 5. Trigger async scoring — does not block the HTTP response.
	go s.scoreAndAct(payload.AssetID)

	return nil
}

// scoreAndAct computes the health score and applies downstream actions:
//   - persists score to assets.health_score
//   - alerts via NotifyService if score < 60
//   - flags asset on-chain via TxManager if score < 40
func (s *VeriflowService) scoreAndAct(assetID string) {
	ctx := context.Background()

	score, err := s.scorer.ComputeScore(ctx, assetID)
	if err != nil {
		s.logger.Error("veriflow_service.scoreAndAct: score failed",
			zap.String("asset_id", assetID),
			zap.Error(err),
		)
		return
	}

	// Persist score.
	asset, err := s.assetRepo.GetByID(ctx, assetID)
	if err != nil {
		s.logger.Error("veriflow_service.scoreAndAct: fetch asset",
			zap.String("asset_id", assetID),
			zap.Error(err),
		)
		return
	}
	asset.HealthScore = score
	if err := s.assetRepo.Update(ctx, asset); err != nil {
		s.logger.Error("veriflow_service.scoreAndAct: persist score",
			zap.String("asset_id", assetID),
			zap.Error(err),
		)
	}

	// Alert if score is below warning threshold.
	if score < veriflow.ScoreAlertThreshold {
		msg := fmt.Sprintf("Asset %s health score dropped to %.1f — curator review required", assetID, score)
		if err := s.notifySvc.Send(ctx, asset.OwnerID, "health_alert", msg, ChannelInApp); err != nil {
			s.logger.Error("veriflow_service.scoreAndAct: notify failed", zap.Error(err))
		}
	}

	// Flag on-chain if score is critically low and asset is still ACTIVE.
	if score < veriflow.ScoreFlagThreshold && asset.Status == models.AssetStatusActive {
		assetID32 := uuidToBytes32(assetID)
		if _, err := s.txManager.UpdateAssetStatusOnChain(ctx, assetID32, onChainFlagged); err != nil {
			s.logger.Error("veriflow_service.scoreAndAct: on-chain flag failed",
				zap.String("asset_id", assetID),
				zap.Error(err),
			)
		} else {
			s.logger.Warn("asset flagged on-chain due to critical health score",
				zap.String("asset_id", assetID),
				zap.Float64("score", score),
			)
		}
	}
}

func (s *VeriflowService) GetHealthScore(ctx context.Context, assetID string) (dto.HealthScoreResponse, error) {
	asset, err := s.assetRepo.GetByID(ctx, assetID)
	if err != nil {
		return dto.HealthScoreResponse{}, fmt.Errorf("veriflow_service.GetHealthScore: %w", err)
	}
	return dto.HealthScoreResponse{
		AssetID:     asset.ID,
		HealthScore: asset.HealthScore,
		Status:      asset.Status,
		ComputedAt:  asset.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func (s *VeriflowService) GetAttestation(ctx context.Context, assetID string) (dto.AttestationResponse, error) {
	att, err := s.attestationRepo.GetLatestByAsset(ctx, assetID)
	if err != nil {
		return dto.AttestationResponse{}, fmt.Errorf("veriflow_service.GetAttestation: %w", err)
	}
	return dto.AttestationResponse{
		AssetID:     att.AssetID,
		HealthScore: float64(att.HealthScore),
		HealthHash:  att.HealthHash,
		XDCTxHash:   att.XDCTxHash,
		Timestamp:   att.AttestedAt.Format(time.RFC3339),
	}, nil
}
