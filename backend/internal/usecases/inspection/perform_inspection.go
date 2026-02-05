package inspection

import (
	"context"
	"workshop-platform/backend/internal/domain"
	"workshop-platform/backend/internal/repositories"
)

type PerformInspectionInput struct {
	EntryID     string
	VehicleID   string
	TemplateID  string
	PerformedBy string
	Notes       string
	ItemResults []domain.InspectionItemResult
}

type PerformInspection struct {
	inspectionRepo repositories.InspectionRepository
}

func NewPerformInspection(inspectionRepo repositories.InspectionRepository) *PerformInspection {
	return &PerformInspection{inspectionRepo: inspectionRepo}
}

func (uc *PerformInspection) Execute(ctx context.Context, input PerformInspectionInput) (*domain.Inspection, error) {
	ins := &domain.Inspection{
		EntryID:     input.EntryID,
		VehicleID:   input.VehicleID,
		TemplateID:  input.TemplateID,
		PerformedBy: input.PerformedBy,
		Status:      "draft",
		Notes:       input.Notes,
	}
	if err := uc.inspectionRepo.Create(ctx, ins); err != nil {
		return nil, err
	}
	for _, r := range input.ItemResults {
		r.InspectionID = ins.ID
		if err := uc.inspectionRepo.SaveItemResult(ctx, &r); err != nil {
			return nil, err
		}
	}
	return ins, nil
}
