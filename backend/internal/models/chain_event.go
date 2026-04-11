package models

import "time"

// ChainEvent stores indexed XDC contract events.
type ChainEvent struct {
	ID          uint64    `gorm:"primaryKey;autoIncrement"`
	EventType   string    `gorm:"type:varchar(50);not null;index"`
	ContractAddress string `gorm:"type:varchar(42);not null"`
	BlockNumber uint64    `gorm:"not null;index"`
	TxHash      string    `gorm:"type:varchar(66);not null;uniqueIndex"`
	LogIndex    uint      `gorm:"not null"`
	ParsedArgs  string    `gorm:"type:json"`
	CreatedAt   time.Time
}

func (ChainEvent) TableName() string {
	return "chain_events"
}
