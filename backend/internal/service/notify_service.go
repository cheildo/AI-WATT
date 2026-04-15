package service

import (
	"context"

	"go.uber.org/zap"
)

// NotifyChannel represents a delivery channel for notifications.
type NotifyChannel string

const (
	ChannelEmail   NotifyChannel = "email"
	ChannelWebhook NotifyChannel = "webhook"
	ChannelInApp   NotifyChannel = "in_app"
)

// NotifyServicer sends alerts for protocol events.
type NotifyServicer interface {
	Send(ctx context.Context, userID, event, message string, channels ...NotifyChannel) error
}

// NotifyService implements NotifyServicer.
// Phase 9 will wire real email + webhook delivery.
type NotifyService struct {
	logger *zap.Logger
}

func NewNotifyService(logger *zap.Logger) *NotifyService {
	return &NotifyService{logger: logger}
}

func (s *NotifyService) Send(ctx context.Context, userID, event, message string, channels ...NotifyChannel) error {
	s.logger.Info("notification",
		zap.String("user_id", userID),
		zap.String("event", event),
		zap.String("message", message),
	)
	return nil
}
