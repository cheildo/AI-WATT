package repository

import (
	"context"

	"gorm.io/gorm"

	"github.com/neurowatt/aiwatt-backend/internal/models"
)

// RepaymentRepo defines the repayment database access interface.
type RepaymentRepo interface {
	Create(ctx context.Context, r *models.Repayment) error
	GetByLoan(ctx context.Context, loanID string) ([]*models.Repayment, error)
	SumByLoan(ctx context.Context, loanID string) (uint64, error)
}

// RepaymentRepository implements RepaymentRepo using GORM.
type RepaymentRepository struct {
	db *gorm.DB
}

func NewRepaymentRepository(db *gorm.DB) *RepaymentRepository {
	return &RepaymentRepository{db: db}
}

func (r *RepaymentRepository) Create(ctx context.Context, rep *models.Repayment) error {
	return r.db.WithContext(ctx).Create(rep).Error
}

func (r *RepaymentRepository) GetByLoan(ctx context.Context, loanID string) ([]*models.Repayment, error) {
	var rows []*models.Repayment
	err := r.db.WithContext(ctx).
		Where("loan_id = ?", loanID).
		Order("paid_at DESC").
		Find(&rows).Error
	return rows, err
}

func (r *RepaymentRepository) SumByLoan(ctx context.Context, loanID string) (uint64, error) {
	var total uint64
	err := r.db.WithContext(ctx).
		Model(&models.Repayment{}).
		Where("loan_id = ?", loanID).
		Select("COALESCE(SUM(amount), 0)").
		Scan(&total).Error
	return total, err
}
