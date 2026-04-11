package veriflow

import (
	"context"
	"fmt"
	"time"

	"go.uber.org/zap"

	"github.com/neurowatt/aiwatt-backend/internal/models"
	"github.com/neurowatt/aiwatt-backend/internal/repository"
	"github.com/neurowatt/aiwatt-backend/pkg/crypto"
)

// Ingestor receives and validates telemetry payloads from Veriflow agents.
type Ingestor struct {
	telemetryRepo repository.TelemetryRepo
	logger        *zap.Logger
	// agentSecrets maps assetID → HMAC key (loaded from DB or config)
	agentSecrets map[string][]byte
}

// NewIngestor constructs an Ingestor.
func NewIngestor(repo repository.TelemetryRepo, logger *zap.Logger) *Ingestor {
	return &Ingestor{
		telemetryRepo: repo,
		logger:        logger,
		agentSecrets:  make(map[string][]byte),
	}
}

// IngestPayload validates the HMAC signature and writes the telemetry row.
func (i *Ingestor) IngestPayload(ctx context.Context, assetID, hmacSig string, payload []byte, reading models.Telemetry) error {
	secret, ok := i.agentSecrets[assetID]
	if !ok {
		return fmt.Errorf("ingestor.IngestPayload: unknown asset %s", assetID)
	}
	if !crypto.HMACVerify(secret, payload, hmacSig) {
		return fmt.Errorf("ingestor.IngestPayload: HMAC verification failed for asset %s", assetID)
	}
	reading.AssetID = assetID
	reading.HMACSignature = hmacSig
	reading.RecordedAt = time.Now().UTC()
	if err := i.telemetryRepo.Create(ctx, &reading); err != nil {
		return fmt.Errorf("ingestor.IngestPayload: persist: %w", err)
	}
	i.logger.Info("telemetry ingested", zap.String("asset_id", assetID))
	return nil
}
