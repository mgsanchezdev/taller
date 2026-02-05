package history

import (
	"context"
	"workshop-platform/backend/internal/domain"
	"workshop-platform/backend/internal/repositories"
)

type SearchByPlateInput struct {
	WorkshopID string
	Plate      string
}

type SearchByPlateResult struct {
	Vehicle  *domain.Vehicle
	Entries  []*domain.WorkshopEntry
}

type SearchByPlate struct {
	vehicleRepo repositories.VehicleRepository
	entryRepo   repositories.EntryRepository
}

func NewSearchByPlate(
	vehicleRepo repositories.VehicleRepository,
	entryRepo repositories.EntryRepository,
) *SearchByPlate {
	return &SearchByPlate{vehicleRepo: vehicleRepo, entryRepo: entryRepo}
}

func (uc *SearchByPlate) Execute(ctx context.Context, input SearchByPlateInput) (*SearchByPlateResult, error) {
	v, err := uc.vehicleRepo.GetByPlate(ctx, input.WorkshopID, input.Plate)
	if err != nil || v == nil {
		return nil, err
	}
	entries, _ := uc.entryRepo.ListByVehicle(ctx, v.ID)
	return &SearchByPlateResult{Vehicle: v, Entries: entries}, nil
}
