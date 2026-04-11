package service

import (
	"context"

	"github.com/neurowatt/aiwatt-backend/internal/api/dto"
)

// VeriflowServicer defines the Veriflow telemetry + scoring interface.
type VeriflowServicer interface {
	IngestTelemetry(ctx context.Context, payload dto.TelemetryPayload) error
	GetHealthScore(ctx context.Context, assetID string) (dto.HealthScoreResponse, error)
	GetAttestation(ctx context.Context, assetID string) (dto.AttestationResponse, error)
}

// VeriflowService implements VeriflowServicer.
type VeriflowService struct {
	// TODO: inject TelemetryRepo, veriflow ingestor + scorer
}

// NewVeriflowService constructs a VeriflowService.
func NewVeriflowService() *VeriflowService {
	return &VeriflowService{}
}

func (s *VeriflowService) IngestTelemetry(ctx context.Context, payload dto.TelemetryPayload) error {
	// TODO: verify HMAC, write telemetry row, trigger scorer
	return nil
}

func (s *VeriflowService) GetHealthScore(ctx context.Context, assetID string) (dto.HealthScoreResponse, error) {
	// TODO: fetch latest score from TelemetryRepo
	return dto.HealthScoreResponse{}, nil
}

func (s *VeriflowService) GetAttestation(ctx context.Context, assetID string) (dto.AttestationResponse, error) {
	// TODO: fetch latest attestation record from DB
	return dto.AttestationResponse{}, nil
}
