package repositories

import (
	"context"
	"workshop-platform/backend/internal/domain"
)

type EntryRepository interface {
	Create(ctx context.Context, e *domain.WorkshopEntry) error
	GetByID(ctx context.Context, id string) (*domain.WorkshopEntry, error)
	GetByEntryNumber(ctx context.Context, workshopID, entryNumber string) (*domain.WorkshopEntry, error)
	ListByWorkshop(ctx context.Context, workshopID string, limit, offset int) ([]*domain.WorkshopEntry, int, error)
	ListByVehicle(ctx context.Context, vehicleID string) ([]*domain.WorkshopEntry, error)
	Update(ctx context.Context, e *domain.WorkshopEntry) error
}
