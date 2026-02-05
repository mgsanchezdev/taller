package repositories

import (
	"context"
	"workshop-platform/backend/internal/domain"
)

type MediaRepository interface {
	Create(ctx context.Context, m *domain.Media) error
	GetByID(ctx context.Context, id string) (*domain.Media, error)
	ListByEntity(ctx context.Context, entityType, entityID string) ([]*domain.Media, error)
	Delete(ctx context.Context, id string) error
}
