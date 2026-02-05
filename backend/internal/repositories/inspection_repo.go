package repositories

import (
	"context"
	"workshop-platform/backend/internal/domain"
)

type InspectionRepository interface {
	Create(ctx context.Context, i *domain.Inspection) error
	GetByID(ctx context.Context, id string) (*domain.Inspection, error)
	ListByEntry(ctx context.Context, entryID string) ([]*domain.Inspection, error)
	ListByVehicle(ctx context.Context, vehicleID string) ([]*domain.Inspection, error)
	Update(ctx context.Context, i *domain.Inspection) error
	SaveItemResult(ctx context.Context, r *domain.InspectionItemResult) error
	GetItemResults(ctx context.Context, inspectionID string) ([]*domain.InspectionItemResult, error)
}
