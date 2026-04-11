package repository

import (
	"context"

	"gorm.io/gorm"

	"github.com/neurowatt/aiwatt-backend/internal/models"
)

// UserRepo defines the user database access interface.
type UserRepo interface {
	Create(ctx context.Context, user *models.User) error
	GetByID(ctx context.Context, id string) (*models.User, error)
	GetByEmail(ctx context.Context, email string) (*models.User, error)
	GetByWallet(ctx context.Context, walletAddress string) (*models.User, error)
	Update(ctx context.Context, user *models.User) error
	List(ctx context.Context, role string, offset, limit int) ([]*models.User, int64, error)
}

// UserRepository implements UserRepo using GORM.
type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository constructs a UserRepository.
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(ctx context.Context, user *models.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}

func (r *UserRepository) GetByID(ctx context.Context, id string) (*models.User, error) {
	var u models.User
	if err := r.db.WithContext(ctx).First(&u, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	var u models.User
	if err := r.db.WithContext(ctx).First(&u, "email = ?", email).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *UserRepository) GetByWallet(ctx context.Context, walletAddress string) (*models.User, error) {
	var u models.User
	if err := r.db.WithContext(ctx).First(&u, "wallet_address = ?", walletAddress).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *UserRepository) Update(ctx context.Context, user *models.User) error {
	return r.db.WithContext(ctx).Save(user).Error
}

func (r *UserRepository) List(ctx context.Context, role string, offset, limit int) ([]*models.User, int64, error) {
	var users []*models.User
	var total int64
	q := r.db.WithContext(ctx).Model(&models.User{})
	if role != "" {
		q = q.Where("role = ?", role)
	}
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err := q.Offset(offset).Limit(limit).Find(&users).Error; err != nil {
		return nil, 0, err
	}
	return users, total, nil
}
