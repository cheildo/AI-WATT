package models

import (
	"time"

	"gorm.io/gorm"
)

// Loan engine types matching the three-engine credit model.
const (
	EnginePreDelivery      = 1
	EnginePostDelivery     = 2
	EngineCapitalReactivation = 3
)

// Loan statuses.
const (
	LoanStatusPending   = "pending"
	LoanStatusActive    = "active"
	LoanStatusRepaying  = "repaying"
	LoanStatusSettled   = "settled"
	LoanStatusDefaulted = "defaulted"
)

// Loan is the GORM model for the loans table.
type Loan struct {
	ID          string         `gorm:"type:char(36);primaryKey"`
	AssetID     string         `gorm:"type:char(36);not null;index"`
	BorrowerID  string         `gorm:"type:char(36);not null;index"`
	CuratorID   string         `gorm:"type:char(36);index"`
	Amount      float64        `gorm:"type:decimal(20,6);not null"`
	EngineType  int            `gorm:"type:tinyint;not null"`
	Status      string         `gorm:"type:varchar(20);not null;default:'pending'"`
	OnChainTxHash string       `gorm:"type:varchar(66)"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func (Loan) TableName() string {
	return "loans"
}
