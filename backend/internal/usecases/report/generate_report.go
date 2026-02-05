package report

import (
	"context"
	"workshop-platform/backend/internal/domain"
	"workshop-platform/backend/internal/repositories"
)

type GenerateReportInput struct {
	WorkshopID string
	EntryID    string
	VehicleID  string
	Type       string // inspection, repair, full
}

type GenerateReport struct {
	reportRepo repositories.ReportRepository
}

func NewGenerateReport(reportRepo repositories.ReportRepository) *GenerateReport {
	return &GenerateReport{reportRepo: reportRepo}
}

func (uc *GenerateReport) Execute(ctx context.Context, input GenerateReportInput) (*domain.Report, error) {
	r := &domain.Report{
		WorkshopID: input.WorkshopID,
		EntryID:    input.EntryID,
		VehicleID:  input.VehicleID,
		Type:       input.Type,
		PDFURL:     "", // filled by PDF generator
	}
	if err := uc.reportRepo.Create(ctx, r); err != nil {
		return nil, err
	}
	return r, nil
}
