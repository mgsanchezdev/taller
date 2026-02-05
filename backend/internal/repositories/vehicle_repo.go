package repositories

import (
	"context"
	"workshop-platform/backend/internal/domain"
)

type VehicleRepository interface {
	Create(ctx context.Context, v *domain.Vehicle) error
	GetByID(ctx context.Context, id string) (*domain.Vehicle, error)
	GetByPlate(ctx context.Context, workshopID, plate string) (*domain.Vehicle, error)
	ListByWorkshop(ctx context.Context, workshopID string, limit, offset int) ([]*domain.Vehicle, int, error)
	Update(ctx context.Context, v *domain.Vehicle) error
	Delete(ctx context.Context, id string) error
}
