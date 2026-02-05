package repositories

import (
	"context"
	"workshop-platform/backend/internal/domain"
)

type ReportRepository interface {
	Create(ctx context.Context, r *domain.Report) error
	GetByID(ctx context.Context, id string) (*domain.Report, error)
	ListByEntry(ctx context.Context, entryID string) ([]*domain.Report, error)
	ListByVehicle(ctx context.Context, vehicleID string) ([]*domain.Report, error)
}
