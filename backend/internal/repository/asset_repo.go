package repository

import (
	"context"

	"gorm.io/gorm"

	"github.com/neurowatt/aiwatt-backend/internal/models"
)

// AssetRepo defines the asset database access interface.
type AssetRepo interface {
	Create(ctx context.Context, asset *models.Asset) error
	GetByID(ctx context.Context, id string) (*models.Asset, error)
	Update(ctx context.Context, asset *models.Asset) error
	List(ctx context.Context, ownerID, status string, offset, limit int) ([]*models.Asset, int64, error)
}

// AssetRepository implements AssetRepo using GORM.
type AssetRepository struct {
	db *gorm.DB
}

func NewAssetRepository(db *gorm.DB) *AssetRepository {
	return &AssetRepository{db: db}
}

func (r *AssetRepository) Create(ctx context.Context, asset *models.Asset) error {
	return r.db.WithContext(ctx).Create(asset).Error
}

func (r *AssetRepository) GetByID(ctx context.Context, id string) (*models.Asset, error) {
	var a models.Asset
	if err := r.db.WithContext(ctx).First(&a, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &a, nil
}

func (r *AssetRepository) Update(ctx context.Context, asset *models.Asset) error {
	return r.db.WithContext(ctx).Save(asset).Error
}

func (r *AssetRepository) List(ctx context.Context, ownerID, status string, offset, limit int) ([]*models.Asset, int64, error) {
	var assets []*models.Asset
	var total int64
	q := r.db.WithContext(ctx).Model(&models.Asset{})
	if ownerID != "" {
		q = q.Where("owner_id = ?", ownerID)
	}
	if status != "" {
		q = q.Where("status = ?", status)
	}
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err := q.Offset(offset).Limit(limit).Find(&assets).Error; err != nil {
		return nil, 0, err
	}
	return assets, total, nil
}
