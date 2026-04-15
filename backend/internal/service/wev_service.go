package service

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/neurowatt/aiwatt-backend/internal/api/dto"
	"github.com/neurowatt/aiwatt-backend/internal/models"
	"github.com/neurowatt/aiwatt-backend/internal/repository"
)

// WEVServicer manages the sWATT redemption queue.
type WEVServicer interface {
	Enqueue(ctx context.Context, userID string, req dto.RedemptionRequest) (dto.RedemptionResponse, error)
	Cancel(ctx context.Context, userID string, requestID string) error
	GetQueueStatus(ctx context.Context) (dto.QueueStatusResponse, error)
	GetUserQueue(ctx context.Context, userID string) ([]dto.RedemptionResponse, error)
}

// WEVService implements WEVServicer.
type WEVService struct {
	wevRepo repository.WEVRepo
}

func NewWEVService(wevRepo repository.WEVRepo) *WEVService {
	return &WEVService{wevRepo: wevRepo}
}

func (s *WEVService) Enqueue(ctx context.Context, userID string, req dto.RedemptionRequest) (dto.RedemptionResponse, error) {
	estimatedDays := 30
	if req.Priority {
		estimatedDays = 3
	}

	entry := &models.WEVQueueEntry{
		ID:          uuid.NewString(),
		RequestID:   uuid.NewString(), // placeholder — real ID comes from contract in Phase 8
		UserID:      userID,
		SWattAmount: req.SWattAmount,
		PriorityFee: req.PriorityFee,
		Status:      models.WEVStatusQueued,
		RequestedAt: time.Now().UTC(),
	}

	if err := s.wevRepo.Create(ctx, entry); err != nil {
		return dto.RedemptionResponse{}, fmt.Errorf("wev_service.Enqueue: %w", err)
	}

	return dto.RedemptionResponse{
		RequestID:     entry.RequestID,
		EstimatedDays: estimatedDays,
		Status:        entry.Status,
	}, nil
}

func (s *WEVService) Cancel(ctx context.Context, userID string, requestID string) error {
	entry, err := s.wevRepo.GetByRequestID(ctx, requestID)
	if err != nil {
		return fmt.Errorf("wev_service.Cancel: request not found: %w", err)
	}
	if entry.UserID != userID {
		return fmt.Errorf("wev_service.Cancel: not request owner")
	}
	if entry.Status != models.WEVStatusQueued {
		return fmt.Errorf("wev_service.Cancel: request is not queued (status: %s)", entry.Status)
	}

	if err := s.wevRepo.UpdateStatus(ctx, entry.ID, models.WEVStatusCancelled, nil); err != nil {
		return fmt.Errorf("wev_service.Cancel: %w", err)
	}
	return nil
}

func (s *WEVService) GetQueueStatus(ctx context.Context) (dto.QueueStatusResponse, error) {
	pending, err := s.wevRepo.GetPendingBatch(ctx, 10_000)
	if err != nil {
		return dto.QueueStatusResponse{}, fmt.Errorf("wev_service.GetQueueStatus: %w", err)
	}

	var depthSWatt uint64
	for _, e := range pending {
		depthSWatt += e.SWattAmount
	}

	return dto.QueueStatusResponse{
		DepthSWatt:     depthSWatt,
		QueueDepth:     int64(len(pending)),
		NextProcessing: time.Now().UTC().Add(24 * time.Hour).Format(time.RFC3339),
		StandardDays:   30,
		PriorityDays:   3,
		PriorityFeeBPS: 50,
	}, nil
}

func (s *WEVService) GetUserQueue(ctx context.Context, userID string) ([]dto.RedemptionResponse, error) {
	entries, err := s.wevRepo.GetByUser(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("wev_service.GetUserQueue: %w", err)
	}

	resp := make([]dto.RedemptionResponse, len(entries))
	for i, e := range entries {
		days := 30
		if e.PriorityFee > 0 {
			days = 3
		}
		resp[i] = dto.RedemptionResponse{
			RequestID:     e.RequestID,
			EstimatedDays: days,
			Status:        e.Status,
		}
	}
	return resp, nil
}
