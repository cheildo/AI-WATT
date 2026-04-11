package veriflow

import (
	"context"
	"fmt"
	"time"

	"go.uber.org/zap"

	"github.com/neurowatt/aiwatt-backend/internal/repository"
	pkgcrypto "github.com/neurowatt/aiwatt-backend/pkg/crypto"
)

// Attester writes daily health attestations to the HealthAttestation.sol contract.
type Attester struct {
	telemetryRepo repository.TelemetryRepo
	scorer        *Scorer
	// TODO: inject TxManager + blockchain client for contract interaction
	logger *zap.Logger
}

// NewAttester constructs an Attester.
func NewAttester(repo repository.TelemetryRepo, scorer *Scorer, logger *zap.Logger) *Attester {
	return &Attester{telemetryRepo: repo, scorer: scorer, logger: logger}
}

// WriteAttestation computes the health hash and submits it to HealthAttestation.sol.
func (a *Attester) WriteAttestation(ctx context.Context, assetID string) error {
	score, _, err := a.scorer.ComputeScore(ctx, assetID)
	if err != nil {
		return fmt.Errorf("attester.WriteAttestation: score: %w", err)
	}

	now := time.Now().UTC()
	raw := fmt.Sprintf("%s:%.2f:%d", assetID, score, now.Unix())
	hash := pkgcrypto.Keccak256Hex([]byte(raw))

	a.logger.Info("attestation prepared",
		zap.String("asset_id", assetID),
		zap.Float64("score", score),
		zap.String("health_hash", hash),
		zap.Time("timestamp", now),
	)
	// TODO: call TxManager.SendTransaction with HealthAttestation.sol submitAttestation() calldata
	// TODO: store attestation pre-image in MySQL (attestations table)
	return nil
}
