package notification

import (
	"context"
	"workshop-platform/backend/internal/domain"
	"workshop-platform/backend/internal/repositories"
)

type NotifyWhatsAppInput struct {
	To   string
	Body string
	EntryID string
	VehicleID string
}

type NotifyWhatsApp struct {
	notificationRepo repositories.NotificationRepository
}

func NewNotifyWhatsApp(notificationRepo repositories.NotificationRepository) *NotifyWhatsApp {
	return &NotifyWhatsApp{notificationRepo: notificationRepo}
}

func (uc *NotifyWhatsApp) Execute(ctx context.Context, input NotifyWhatsAppInput) (*domain.Notification, error) {
	n := &domain.Notification{
		Channel:   "whatsapp",
		To:        input.To,
		Body:      input.Body,
		EntryID:   input.EntryID,
		VehicleID: input.VehicleID,
		Status:    "pending",
	}
	if err := uc.notificationRepo.Create(ctx, n); err != nil {
		return nil, err
	}
	// TODO: call WhatsApp API and update n.Status / n.SentAt
	return n, nil
}
