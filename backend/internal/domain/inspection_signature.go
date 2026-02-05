package domain

import "time"

type InspectionSignature struct {
	ID           string    `json:"id"`
	InspectionID string    `json:"inspection_id"`
	SignerRole   string    `json:"signer_role"` // client, technician, etc.
	SignerName   string    `json:"signer_name"`
	SignatureURL string    `json:"signature_url"`
	SignedAt     time.Time `json:"signed_at"`
	IP           string    `json:"ip,omitempty"`
	CreatedAt    time.Time `json:"created_at"`
}
