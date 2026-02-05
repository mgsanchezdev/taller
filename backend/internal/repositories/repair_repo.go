package repositories

import (
	"context"
	"workshop-platform/backend/internal/domain"
)

type RepairRepository interface {
	Create(ctx context.Context, r *domain.Repair) error
	GetByID(ctx context.Context, id string) (*domain.Repair, error)
	ListByEntry(ctx context.Context, entryID string) ([]*domain.Repair, error)
	ListByVehicle(ctx context.Context, vehicleID string) ([]*domain.Repair, error)
	Update(ctx context.Context, r *domain.Repair) error
}
