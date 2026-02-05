package repositories

import (
	"context"
	"workshop-platform/backend/internal/domain"
)

type UserRepository interface {
	Create(ctx context.Context, u *domain.User) error
	GetByID(ctx context.Context, id string) (*domain.User, error)
	GetByEmail(ctx context.Context, workshopID, email string) (*domain.User, error)
	ListByWorkshop(ctx context.Context, workshopID string) ([]*domain.User, error)
	Update(ctx context.Context, u *domain.User) error
}
