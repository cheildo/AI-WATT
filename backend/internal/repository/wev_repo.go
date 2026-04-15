package repository

import (
	"context"
	"time"

	"gorm.io/gorm"

	"github.com/neurowatt/aiwatt-backend/internal/models"
)

// WEVRepo defines the WEV queue database access interface.
type WEVRepo interface {
	Create(ctx context.Context, e *models.WEVQueueEntry) error
	GetByRequestID(ctx context.Context, requestID string) (*models.WEVQueueEntry, error)
	GetByUser(ctx context.Context, userID string) ([]*models.WEVQueueEntry, error)
	UpdateStatus(ctx context.Context, id string, status string, processedAt *time.Time) error
	GetPendingBatch(ctx context.Context, limit int) ([]*models.WEVQueueEntry, error)
}

// WEVRepository implements WEVRepo using GORM.
type WEVRepository struct {
	db *gorm.DB
}

func NewWEVRepository(db *gorm.DB) *WEVRepository {
	return &WEVRepository{db: db}
}

func (r *WEVRepository) Create(ctx context.Context, e *models.WEVQueueEntry) error {
	return r.db.WithContext(ctx).Create(e).Error
}

func (r *WEVRepository) GetByRequestID(ctx context.Context, requestID string) (*models.WEVQueueEntry, error) {
	var e models.WEVQueueEntry
	err := r.db.WithContext(ctx).Where("request_id = ?", requestID).First(&e).Error
	if err != nil {
		return nil, err
	}
	return &e, nil
}

func (r *WEVRepository) GetByUser(ctx context.Context, userID string) ([]*models.WEVQueueEntry, error) {
	var rows []*models.WEVQueueEntry
	err := r.db.WithContext(ctx).
		Where("user_id = ?", userID).
		Order("requested_at DESC").
		Find(&rows).Error
	return rows, err
}

func (r *WEVRepository) UpdateStatus(ctx context.Context, id string, status string, processedAt *time.Time) error {
	updates := map[string]any{"status": status}
	if processedAt != nil {
		updates["processed_at"] = processedAt
	}
	return r.db.WithContext(ctx).
		Model(&models.WEVQueueEntry{}).
		Where("id = ?", id).
		Updates(updates).Error
}

func (r *WEVRepository) GetPendingBatch(ctx context.Context, limit int) ([]*models.WEVQueueEntry, error) {
	var rows []*models.WEVQueueEntry
	err := r.db.WithContext(ctx).
		Where("status = ?", models.WEVStatusQueued).
		Order("priority_fee DESC, requested_at ASC"). // priority first, then FIFO
		Limit(limit).
		Find(&rows).Error
	return rows, err
}
