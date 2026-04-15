package models

import "time"

// Repayment records a single WATT repayment event for a loan.
type Repayment struct {
	ID      string    `gorm:"type:char(36);primaryKey"`
	LoanID  string    `gorm:"type:varchar(66);not null;index"` // bytes32 hex
	Amount  uint64    `gorm:"type:bigint unsigned;not null"`   // WATT, 6 decimals
	TxHash  string    `gorm:"type:varchar(66)"`
	PaidAt  time.Time `gorm:"not null"`
}

func (Repayment) TableName() string { return "repayments" }
