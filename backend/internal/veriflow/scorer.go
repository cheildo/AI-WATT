package veriflow

import (
	"context"
	"fmt"

	"go.uber.org/zap"

	"github.com/neurowatt/aiwatt-backend/internal/models"
	"github.com/neurowatt/aiwatt-backend/internal/repository"
)

const (
	scoreThreshold        = 60.0
	minGPUUtilization     = 70.0
	maxGPUTemperature     = 85.0
	maxGPUErrorRate       = 0.0001
	heartbeatTimeoutMins  = 15
)

// Scorer computes the health score (0–100) for a hardware asset using rule-based logic.
type Scorer struct {
	telemetryRepo repository.TelemetryRepo
	logger        *zap.Logger
}

// NewScorer constructs a Scorer.
func NewScorer(repo repository.TelemetryRepo, logger *zap.Logger) *Scorer {
	return &Scorer{telemetryRepo: repo, logger: logger}
}

// ComputeScore calculates a health score for assetID from its latest telemetry.
// Returns score 0–100 and whether the asset should be flagged.
func (s *Scorer) ComputeScore(ctx context.Context, assetID string) (float64, bool, error) {
	t, err := s.telemetryRepo.GetLatestByAsset(ctx, assetID)
	if err != nil {
		return 0, false, fmt.Errorf("scorer.ComputeScore: fetch telemetry: %w", err)
	}
	score := s.score(t)
	flagged := score < scoreThreshold
	s.logger.Info("health score computed",
		zap.String("asset_id", assetID),
		zap.Float64("score", score),
		zap.Bool("flagged", flagged),
	)
	return score, flagged, nil
}

func (s *Scorer) score(t *models.Telemetry) float64 {
	var score float64 = 100

	if t.GPUUtilization < minGPUUtilization {
		score -= 25
	}
	if t.GPUTemperature > maxGPUTemperature {
		score -= 25
	}
	if t.GPUErrorRate > maxGPUErrorRate {
		score -= 30
	}
	// Additional deductions for anomalous readings can be added here
	if score < 0 {
		score = 0
	}
	return score
}
