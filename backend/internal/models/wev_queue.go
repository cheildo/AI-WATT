package models

import "time"

// WEV queue statuses.
const (
	WEVStatusQueued     = "queued"
	WEVStatusProcessing = "processing"
	WEVStatusFulfilled  = "fulfilled"
	WEVStatusCancelled  = "cancelled"
)

// WEVQueueEntry tracks a sWATT redemption request in the off-chain DB.
type WEVQueueEntry struct {
	ID          string     `gorm:"type:char(36);primaryKey"`
	RequestID   string     `gorm:"type:varchar(66);uniqueIndex;not null"` // bytes32 hex from contract
	UserID      string     `gorm:"type:char(36);not null;index"`
	SWattAmount uint64     `gorm:"type:bigint unsigned;not null"` // sWATT shares, 6 decimals
	PriorityFee uint64     `gorm:"type:bigint unsigned;not null;default:0"`
	Status      string     `gorm:"type:varchar(20);not null;default:'queued'"`
	RequestedAt time.Time  `gorm:"not null"`
	ProcessedAt *time.Time `gorm:"default:null"`
}

func (WEVQueueEntry) TableName() string { return "wev_queue" }
