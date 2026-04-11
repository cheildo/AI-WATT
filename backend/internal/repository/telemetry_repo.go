package repository

import (
	"context"
	"time"

	"gorm.io/gorm"

	"github.com/neurowatt/aiwatt-backend/internal/models"
)

// TelemetryRepo defines the telemetry database access interface.
type TelemetryRepo interface {
	Create(ctx context.Context, t *models.Telemetry) error
	GetLatestByAsset(ctx context.Context, assetID string) (*models.Telemetry, error)
	ListByAssetAndDateRange(ctx context.Context, assetID string, from, to time.Time, offset, limit int) ([]*models.Telemetry, error)
}

// TelemetryRepository implements TelemetryRepo using GORM.
type TelemetryRepository struct {
	db *gorm.DB
}

func NewTelemetryRepository(db *gorm.DB) *TelemetryRepository {
	return &TelemetryRepository{db: db}
}

func (r *TelemetryRepository) Create(ctx context.Context, t *models.Telemetry) error {
	return r.db.WithContext(ctx).Create(t).Error
}

func (r *TelemetryRepository) GetLatestByAsset(ctx context.Context, assetID string) (*models.Telemetry, error) {
	var t models.Telemetry
	if err := r.db.WithContext(ctx).
		Where("asset_id = ?", assetID).
		Order("recorded_at DESC").
		First(&t).Error; err != nil {
		return nil, err
	}
	return &t, nil
}

func (r *TelemetryRepository) ListByAssetAndDateRange(ctx context.Context, assetID string, from, to time.Time, offset, limit int) ([]*models.Telemetry, error) {
	var rows []*models.Telemetry
	if err := r.db.WithContext(ctx).
		Where("asset_id = ? AND recorded_at BETWEEN ? AND ?", assetID, from, to).
		Order("recorded_at DESC").
		Offset(offset).Limit(limit).
		Find(&rows).Error; err != nil {
		return nil, err
	}
	return rows, nil
}
