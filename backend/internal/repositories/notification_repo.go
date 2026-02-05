package repositories

import (
	"context"
	"workshop-platform/backend/internal/domain"
)

type NotificationRepository interface {
	Create(ctx context.Context, n *domain.Notification) error
	GetByID(ctx context.Context, id string) (*domain.Notification, error)
	Update(ctx context.Context, n *domain.Notification) error
}
