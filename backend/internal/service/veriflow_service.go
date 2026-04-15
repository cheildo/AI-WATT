package service

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/neurowatt/aiwatt-backend/internal/api/dto"
	"github.com/neurowatt/aiwatt-backend/internal/models"
	"github.com/neurowatt/aiwatt-backend/internal/repository"
	"github.com/neurowatt/aiwatt-backend/pkg/crypto"
)

// VeriflowServicer defines the Veriflow telemetry + scoring interface.
type VeriflowServicer interface {
	IngestTelemetry(ctx context.Context, payload dto.TelemetryPayload, hmacSecret string) error
	GetHealthScore(ctx context.Context, assetID string) (dto.HealthScoreResponse, error)
	GetAttestation(ctx context.Context, assetID string) (dto.AttestationResponse, error)
}

// VeriflowService implements VeriflowServicer.
type VeriflowService struct {
	telemetryRepo  repository.TelemetryRepo
	attestationRepo repository.AttestationRepo
	assetRepo      repository.AssetRepo
}

func NewVeriflowService(
	telemetryRepo repository.TelemetryRepo,
	attestationRepo repository.AttestationRepo,
	assetRepo repository.AssetRepo,
) *VeriflowService {
	return &VeriflowService{
		telemetryRepo:   telemetryRepo,
		attestationRepo: attestationRepo,
		assetRepo:       assetRepo,
	}
}

// IngestTelemetry verifies the HMAC signature, stores the telemetry row, and
// triggers an async health score computation. The hmacSecret is the per-asset
// shared secret provisioned at onboarding and stored in assets.hmac_secret.
func (s *VeriflowService) IngestTelemetry(ctx context.Context, payload dto.TelemetryPayload, hmacSecret string) error {
	// Verify HMAC: sign the payload JSON (without the signature field) and compare.
	payloadBytes, err := json.Marshal(dto.TelemetryPayload{
		AssetID:         payload.AssetID,
		Timestamp:       payload.Timestamp,
		GPUUtilization:  payload.GPUUtilization,
		GPUTemperature:  payload.GPUTemperature,
		GPUMemoryUsedMB: payload.GPUMemoryUsedMB,
		GPUErrorRate:    payload.GPUErrorRate,
		PowerDrawWatts:  payload.PowerDrawWatts,
		FanSpeedRPM:     payload.FanSpeedRPM,
	})
	if err != nil {
		return fmt.Errorf("veriflow_service.IngestTelemetry: marshal: %w", err)
	}
	if !crypto.HMACVerify([]byte(hmacSecret), payloadBytes, payload.HMACSignature) {
		return fmt.Errorf("veriflow_service.IngestTelemetry: hmac mismatch")
	}

	row := &models.Telemetry{
		AssetID:         payload.AssetID,
		GPUUtilization:  payload.GPUUtilization,
		GPUTemperature:  payload.GPUTemperature,
		GPUMemoryUsedMB: payload.GPUMemoryUsedMB,
		GPUErrorRate:    payload.GPUErrorRate,
		PowerDrawWatts:  payload.PowerDrawWatts,
		FanSpeedRPM:     payload.FanSpeedRPM,
		HMACSignature:   payload.HMACSignature,
		RecordedAt:      time.Unix(payload.Timestamp, 0).UTC(),
	}

	if err := s.telemetryRepo.Create(ctx, row); err != nil {
		return fmt.Errorf("veriflow_service.IngestTelemetry: insert: %w", err)
	}
	return nil
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
