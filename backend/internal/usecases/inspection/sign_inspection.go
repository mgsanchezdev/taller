package inspection

import (
	"context"
	"workshop-platform/backend/internal/domain"
	"workshop-platform/backend/internal/repositories"
)

type SignInspectionInput struct {
	InspectionID string
	SignerRole   string
	SignerName   string
	SignatureURL string
	IP           string
}

type SignInspection struct {
	inspectionRepo repositories.InspectionRepository
	signatureRepo  repositories.SignatureRepository
}

func NewSignInspection(
	inspectionRepo repositories.InspectionRepository,
	signatureRepo repositories.SignatureRepository,
) *SignInspection {
	return &SignInspection{inspectionRepo: inspectionRepo, signatureRepo: signatureRepo}
}

func (uc *SignInspection) Execute(ctx context.Context, input SignInspectionInput) error {
	sig := &domain.InspectionSignature{
		InspectionID: input.InspectionID,
		SignerRole:   input.SignerRole,
		SignerName:   input.SignerName,
		SignatureURL: input.SignatureURL,
		IP:           input.IP,
	}
	if err := uc.signatureRepo.Create(ctx, sig); err != nil {
		return err
	}
	ins, _ := uc.inspectionRepo.GetByID(ctx, input.InspectionID)
	if ins != nil {
		ins.Status = "signed"
		_ = uc.inspectionRepo.Update(ctx, ins)
	}
	return nil
}
