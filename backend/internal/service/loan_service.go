package service

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/neurowatt/aiwatt-backend/internal/api/dto"
	"github.com/neurowatt/aiwatt-backend/internal/blockchain"
	"github.com/neurowatt/aiwatt-backend/internal/models"
	"github.com/neurowatt/aiwatt-backend/internal/repository"
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
	loanRepo  repository.LoanRepo
	assetRepo repository.AssetRepo
	bcClient  *blockchain.BlockchainClient
}

func NewLoanService(loanRepo repository.LoanRepo, assetRepo repository.AssetRepo, bcClient *blockchain.BlockchainClient) *LoanService {
	return &LoanService{loanRepo: loanRepo, assetRepo: assetRepo, bcClient: bcClient}
}

// Create creates a loan application after verifying the asset is ACTIVE on-chain.
func (s *LoanService) Create(ctx context.Context, req dto.CreateLoanRequest) (dto.LoanResponse, error) {
	// Verify the asset exists in the DB.
	asset, err := s.assetRepo.GetByID(ctx, req.AssetID)
	if err != nil {
		return dto.LoanResponse{}, fmt.Errorf("loan_service.Create: asset not found: %w", err)
	}

	// Verify asset is ACTIVE in AssetRegistry on XDC.
	assetID32 := uuidToBytes32(asset.ID)
	active, err := s.bcClient.IsAssetActive(ctx, assetID32)
	if err != nil {
		return dto.LoanResponse{}, fmt.Errorf("loan_service.Create: chain check failed: %w", err)
	}
	if !active {
		return dto.LoanResponse{}, fmt.Errorf("loan_service.Create: asset is not ACTIVE on-chain")
	}

	loan := &models.Loan{
		ID:         uuid.NewString(),
		AssetID:    req.AssetID,
		BorrowerID: req.BorrowerID,
		EngineType: req.EngineType,
		Status:     models.LoanStatusPending,
	}
	if err := s.loanRepo.Create(ctx, loan); err != nil {
		return dto.LoanResponse{}, fmt.Errorf("loan_service.Create: %w", err)
	}
	return toLoanDTO(loan), nil
}

func (s *LoanService) GetByID(ctx context.Context, id string) (dto.LoanResponse, error) {
	loan, err := s.loanRepo.GetByID(ctx, id)
	if err != nil {
		return dto.LoanResponse{}, fmt.Errorf("loan_service.GetByID: %w", err)
	}
	return toLoanDTO(loan), nil
}

func (s *LoanService) Update(ctx context.Context, id string, req dto.UpdateLoanRequest) (dto.LoanResponse, error) {
	loan, err := s.loanRepo.GetByID(ctx, id)
	if err != nil {
		return dto.LoanResponse{}, fmt.Errorf("loan_service.Update: %w", err)
	}
	loan.Status = req.Status
	if req.CuratorID != "" {
		loan.CuratorID = req.CuratorID
	}
	if err := s.loanRepo.Update(ctx, loan); err != nil {
		return dto.LoanResponse{}, fmt.Errorf("loan_service.Update: save: %w", err)
	}
	return toLoanDTO(loan), nil
}

func (s *LoanService) List(ctx context.Context, q dto.ListLoansQuery) (dto.ListLoansResponse, error) {
	offset := (q.Page - 1) * q.PageSize
	loans, total, err := s.loanRepo.List(ctx, q.Status, q.BorrowerID, offset, q.PageSize)
	if err != nil {
		return dto.ListLoansResponse{}, fmt.Errorf("loan_service.List: %w", err)
	}
	resp := make([]dto.LoanResponse, len(loans))
	for i, l := range loans {
		resp[i] = toLoanDTO(l)
	}
	return dto.ListLoansResponse{Loans: resp, Total: total, Page: q.Page}, nil
}

func toLoanDTO(l *models.Loan) dto.LoanResponse {
	return dto.LoanResponse{
		ID:            l.ID,
		AssetID:       l.AssetID,
		BorrowerID:    l.BorrowerID,
		CuratorID:     l.CuratorID,
		EngineType:    l.EngineType,
		Status:        l.Status,
		OnChainTxHash: l.OnChainTxHash,
		CreatedAt:     l.CreatedAt.Format(time.RFC3339),
	}
}
