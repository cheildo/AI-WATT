package repository

import (
	"context"

	"gorm.io/gorm"

	"github.com/neurowatt/aiwatt-backend/internal/models"
)

// EventRepo defines the chain event database access interface.
type EventRepo interface {
	Create(ctx context.Context, event *models.ChainEvent) error
	GetByTxHash(ctx context.Context, txHash string) (*models.ChainEvent, error)
	ListByEventType(ctx context.Context, eventType string, offset, limit int) ([]*models.ChainEvent, int64, error)
	GetLatestBlock(ctx context.Context) (uint64, error)
}

// EventRepository implements EventRepo using GORM.
type EventRepository struct {
	db *gorm.DB
}

func NewEventRepository(db *gorm.DB) *EventRepository {
	return &EventRepository{db: db}
}

func (r *EventRepository) Create(ctx context.Context, event *models.ChainEvent) error {
	return r.db.WithContext(ctx).Create(event).Error
}

func (r *EventRepository) GetByTxHash(ctx context.Context, txHash string) (*models.ChainEvent, error) {
	var e models.ChainEvent
	if err := r.db.WithContext(ctx).First(&e, "tx_hash = ?", txHash).Error; err != nil {
		return nil, err
	}
	return &e, nil
}

func (r *EventRepository) ListByEventType(ctx context.Context, eventType string, offset, limit int) ([]*models.ChainEvent, int64, error) {
	var events []*models.ChainEvent
	var total int64
	q := r.db.WithContext(ctx).Model(&models.ChainEvent{}).Where("event_type = ?", eventType)
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err := q.Order("block_number DESC").Offset(offset).Limit(limit).Find(&events).Error; err != nil {
		return nil, 0, err
	}
	return events, total, nil
}

func (r *EventRepository) GetLatestBlock(ctx context.Context) (uint64, error) {
	var event models.ChainEvent
	if err := r.db.WithContext(ctx).Order("block_number DESC").First(&event).Error; err != nil {
		return 0, err
	}
	return event.BlockNumber, nil
}
