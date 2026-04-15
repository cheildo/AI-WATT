package repository

import (
	"context"

	"gorm.io/gorm"

	"github.com/neurowatt/aiwatt-backend/internal/models"
)

// AttestationRepo defines the attestation database access interface.
type AttestationRepo interface {
	Create(ctx context.Context, a *models.Attestation) error
	GetLatestByAsset(ctx context.Context, assetID string) (*models.Attestation, error)
	GetHistory(ctx context.Context, assetID string, limit int) ([]*models.Attestation, error)
}

// AttestationRepository implements AttestationRepo using GORM.
type AttestationRepository struct {
	db *gorm.DB
}

func NewAttestationRepository(db *gorm.DB) *AttestationRepository {
	return &AttestationRepository{db: db}
}

func (r *AttestationRepository) Create(ctx context.Context, a *models.Attestation) error {
	return r.db.WithContext(ctx).Create(a).Error
}

func (r *AttestationRepository) GetLatestByAsset(ctx context.Context, assetID string) (*models.Attestation, error) {
	var a models.Attestation
	err := r.db.WithContext(ctx).
		Where("asset_id = ?", assetID).
		Order("attested_at DESC").
		First(&a).Error
	if err != nil {
		return nil, err
	}
	return &a, nil
}

func (r *AttestationRepository) GetHistory(ctx context.Context, assetID string, limit int) ([]*models.Attestation, error) {
	var rows []*models.Attestation
	err := r.db.WithContext(ctx).
		Where("asset_id = ?", assetID).
		Order("attested_at DESC").
		Limit(limit).
		Find(&rows).Error
	return rows, err
}
