package repair

import (
	"context"
	"time"
	"workshop-platform/backend/internal/domain"
	"workshop-platform/backend/internal/repositories"
)

type RepairVehicleInput struct {
	EntryID      string
	VehicleID    string
	FaultID      string
	Description  string
	LaborHours   float64
	PartsCost    float64
	TechnicianID string
}

type RepairVehicle struct {
	repairRepo repositories.RepairRepository
}

func NewRepairVehicle(repairRepo repositories.RepairRepository) *RepairVehicle {
	return &RepairVehicle{repairRepo: repairRepo}
}

func (uc *RepairVehicle) Execute(ctx context.Context, input RepairVehicleInput) (*domain.Repair, error) {
	r := &domain.Repair{
		EntryID:       input.EntryID,
		VehicleID:     input.VehicleID,
		FaultID:       input.FaultID,
		Description:   input.Description,
		LaborHours:    input.LaborHours,
		PartsCost:     input.PartsCost,
		Status:        "pending",
		TechnicianID:  input.TechnicianID,
	}
	if err := uc.repairRepo.Create(ctx, r); err != nil {
		return nil, err
	}
	return r, nil
}

func (uc *RepairVehicle) Complete(ctx context.Context, repairID string) error {
	r, err := uc.repairRepo.GetByID(ctx, repairID)
	if err != nil || r == nil {
		return err
	}
	now := time.Now()
	r.Status = "completed"
	r.CompletedAt = &now
	return uc.repairRepo.Update(ctx, r)
}
