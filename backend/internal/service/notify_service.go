package service

import "context"

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
type NotifyService struct {
	// TODO: inject email client, webhook client
}

func NewNotifyService() *NotifyService { return &NotifyService{} }

func (s *NotifyService) Send(ctx context.Context, userID, event, message string, channels ...NotifyChannel) error {
	// TODO: fan out to requested delivery channels
	return nil
}
