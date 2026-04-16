package veriflow

import (
	"context"
	"fmt"
	"time"

	"go.uber.org/zap"

	"github.com/neurowatt/aiwatt-backend/internal/repository"
)

const (
	scoringWindow         = 12             // last N telemetry rows (~60 min at 5-min intervals)
	heartbeatTimeoutMins  = 15 * time.Minute

	// GPU utilization thresholds (%)
	utilHigh = 70.0
	utilMid  = 40.0

	// Temperature thresholds (°C)
	tempLow  = 75.0
	tempHigh = 85.0

	// ECC error rate thresholds (fraction, where errCount/10000 = rate)
	eccLow  = 0.0001 // 0.01%
	eccHigh = 0.001  // 0.1%

	// Uptime thresholds (%)
	uptimeHigh = 99.5
	uptimeMid  = 95.0

	// Score thresholds
	ScoreAlertThreshold = 60.0 // below this -> health alert
	ScoreFlagThreshold  = 40.0 // below this AND asset ACTIVE -> flag on-chain
)

// Scorer computes a 0–100 health score for a hardware asset using rule-based logic
// over a rolling window of the last 12 telemetry rows (~60 minutes).
type Scorer struct {
	telemetryRepo repository.TelemetryRepo
	logger        *zap.Logger
}

// NewScorer constructs a Scorer.
func NewScorer(repo repository.TelemetryRepo, logger *zap.Logger) *Scorer {
	return &Scorer{telemetryRepo: repo, logger: logger}
}

// ComputeScore calculates the health score for assetID.
// Returns 0 immediately if the last telemetry is older than heartbeatTimeoutMins.
func (s *Scorer) ComputeScore(ctx context.Context, assetID string) (float64, error) {
	rows, err := s.telemetryRepo.GetLastN(ctx, assetID, scoringWindow)
	if err != nil {
		return 0, fmt.Errorf("scorer.ComputeScore: fetch: %w", err)
	}
	if len(rows) == 0 {
		return 0, fmt.Errorf("scorer.ComputeScore: no telemetry for asset %s", assetID)
	}

	// Heartbeat check — overrides all other scoring.
	if time.Since(rows[0].RecordedAt) > heartbeatTimeoutMins {
		s.logger.Warn("heartbeat timeout — forcing score to 0",
			zap.String("asset_id", assetID),
			zap.Time("last_seen", rows[0].RecordedAt),
		)
		return 0, nil
	}

	// Compute averages over the window.
	var sumUtil, sumTemp, sumECC, sumUptime float64
	for _, r := range rows {
		sumUtil += r.GPUUtilization
		sumTemp += r.GPUTemperature
		// ECCErrors is the raw uncorrected ECC count; normalise to a rate fraction.
		sumECC += float64(r.ECCErrors) / 10000.0
		sumUptime += r.UptimePct
	}
	n := float64(len(rows))
	avgUtil := sumUtil / n
	avgTemp := sumTemp / n
	avgECC := sumECC / n
	avgUptime := sumUptime / n

	var score float64

	// Rule 1: GPU utilisation (25 pts)
	switch {
	case avgUtil >= utilHigh:
		score += 25
	case avgUtil >= utilMid:
		score += 15
	}

	// Rule 2: GPU temperature (25 pts)
	switch {
	case avgTemp <= tempLow:
		score += 25
	case avgTemp <= tempHigh:
		score += 15
	}

	// Rule 3: ECC error rate (25 pts)
	switch {
	case avgECC <= eccLow:
		score += 25
	case avgECC <= eccHigh:
		score += 10
	}

	// Rule 4: Uptime (25 pts)
	switch {
	case avgUptime >= uptimeHigh:
		score += 25
	case avgUptime >= uptimeMid:
		score += 15
	}

	s.logger.Info("health score computed",
		zap.String("asset_id", assetID),
		zap.Float64("score", score),
		zap.Float64("avg_util", avgUtil),
		zap.Float64("avg_temp", avgTemp),
		zap.Float64("avg_ecc", avgECC),
		zap.Float64("avg_uptime", avgUptime),
		zap.Int("rows_sampled", len(rows)),
	)
	return score, nil
}
