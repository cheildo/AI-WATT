package repository

import (
	"context"

	"gorm.io/gorm"

	"github.com/neurowatt/aiwatt-backend/internal/models"
)

// LoanRepo defines the loan database access interface.
type LoanRepo interface {
	Create(ctx context.Context, loan *models.Loan) error
	GetByID(ctx context.Context, id string) (*models.Loan, error)
	Update(ctx context.Context, loan *models.Loan) error
	List(ctx context.Context, status, borrowerID string, offset, limit int) ([]*models.Loan, int64, error)
}

// LoanRepository implements LoanRepo using GORM.
type LoanRepository struct {
	db *gorm.DB
}

func NewLoanRepository(db *gorm.DB) *LoanRepository {
	return &LoanRepository{db: db}
}

func (r *LoanRepository) Create(ctx context.Context, loan *models.Loan) error {
	return r.db.WithContext(ctx).Create(loan).Error
}

func (r *LoanRepository) GetByID(ctx context.Context, id string) (*models.Loan, error) {
	var l models.Loan
	if err := r.db.WithContext(ctx).First(&l, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &l, nil
}

func (r *LoanRepository) Update(ctx context.Context, loan *models.Loan) error {
	return r.db.WithContext(ctx).Save(loan).Error
}

func (r *LoanRepository) List(ctx context.Context, status, borrowerID string, offset, limit int) ([]*models.Loan, int64, error) {
	var loans []*models.Loan
	var total int64
	q := r.db.WithContext(ctx).Model(&models.Loan{})
	if status != "" {
		q = q.Where("status = ?", status)
	}
	if borrowerID != "" {
		q = q.Where("borrower_id = ?", borrowerID)
	}
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err := q.Offset(offset).Limit(limit).Find(&loans).Error; err != nil {
		return nil, 0, err
	}
	return loans, total, nil
}
