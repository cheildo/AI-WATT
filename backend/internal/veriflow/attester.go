package veriflow

import (
	"context"
	"fmt"
	"time"

	"go.uber.org/zap"

	pkgcrypto "github.com/neurowatt/aiwatt-backend/pkg/crypto"
	"github.com/neurowatt/aiwatt-backend/internal/repository"
)

// Attester writes daily health attestations to the HealthAttestation.sol contract.
// Full implementation in Phase 11 — Phase 9 wires the TxManager stub.
type Attester struct {
	telemetryRepo repository.TelemetryRepo
	scorer        *Scorer
	logger        *zap.Logger
}

// NewAttester constructs an Attester.
func NewAttester(repo repository.TelemetryRepo, scorer *Scorer, logger *zap.Logger) *Attester {
	return &Attester{telemetryRepo: repo, scorer: scorer, logger: logger}
}

// WriteAttestation computes the daily health hash and prepares the attestation.
// On-chain submission is completed in Phase 11 when the scheduler is wired.
func (a *Attester) WriteAttestation(ctx context.Context, assetID string) error {
	score, err := a.scorer.ComputeScore(ctx, assetID)
	if err != nil {
		return fmt.Errorf("attester.WriteAttestation: score: %w", err)
	}

	now := time.Now().UTC()
	raw := fmt.Sprintf("%s:%.2f:%d", assetID, score, now.Unix())
	hash := pkgcrypto.Keccak256Hex([]byte(raw))

	a.logger.Info("attestation prepared — on-chain write deferred to Phase 11",
		zap.String("asset_id", assetID),
		zap.Float64("score", score),
		zap.String("health_hash", hash),
		zap.Time("timestamp", now),
	)
	return nil
}
