package domain

import "time"

type Report struct {
	ID          string    `json:"id"`
	WorkshopID  string    `json:"workshop_id"`
	EntryID     string    `json:"entry_id"`
	VehicleID   string    `json:"vehicle_id"`
	Type        string    `json:"type"` // inspection, repair, full
	PDFURL      string    `json:"pdf_url"`
	GeneratedAt time.Time `json:"generated_at"`
	CreatedAt   time.Time `json:"created_at"`
}
