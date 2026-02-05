package domain

import "time"

type Inspection struct {
	ID          string    `json:"id"`
	EntryID     string    `json:"entry_id"`
	VehicleID   string    `json:"vehicle_id"`
	TemplateID  string    `json:"template_id"`
	PerformedBy string    `json:"performed_by"`
	PerformedAt time.Time `json:"performed_at"`
	Status      string    `json:"status"` // draft, completed, signed
	Notes       string    `json:"notes"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
