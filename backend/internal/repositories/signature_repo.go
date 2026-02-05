package repositories

import (
	"context"
	"workshop-platform/backend/internal/domain"
)

type SignatureRepository interface {
	Create(ctx context.Context, s *domain.InspectionSignature) error
	GetByInspection(ctx context.Context, inspectionID string) ([]*domain.InspectionSignature, error)
}
