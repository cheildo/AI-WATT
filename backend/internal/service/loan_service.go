package service

import (
	"context"

	"github.com/neurowatt/aiwatt-backend/internal/api/dto"
)

// LoanServicer defines the loan business logic interface.
type LoanServicer interface {
	Create(ctx context.Context, req dto.CreateLoanRequest) (dto.LoanResponse, error)
	GetByID(ctx context.Context, id string) (dto.LoanResponse, error)
	Update(ctx context.Context, id string, req dto.UpdateLoanRequest) (dto.LoanResponse, error)
	List(ctx context.Context, q dto.ListLoansQuery) (dto.ListLoansResponse, error)
}

// LoanService implements LoanServicer.
type LoanService struct {
	// TODO: inject LoanRepo, AssetService, NotifyService
}

// NewLoanService constructs a LoanService.
func NewLoanService() *LoanService {
	return &LoanService{}
}

func (s *LoanService) Create(ctx context.Context, req dto.CreateLoanRequest) (dto.LoanResponse, error) {
	// TODO: validate asset exists, create loan record, emit notification
	return dto.LoanResponse{}, nil
}

func (s *LoanService) GetByID(ctx context.Context, id string) (dto.LoanResponse, error) {
	// TODO: fetch from LoanRepo, map model → dto
	return dto.LoanResponse{}, nil
}

func (s *LoanService) Update(ctx context.Context, id string, req dto.UpdateLoanRequest) (dto.LoanResponse, error) {
	// TODO: update loan status, trigger downstream effects on status change
	return dto.LoanResponse{}, nil
}

func (s *LoanService) List(ctx context.Context, q dto.ListLoansQuery) (dto.ListLoansResponse, error) {
	// TODO: paginated fetch from LoanRepo
	return dto.ListLoansResponse{}, nil
}
