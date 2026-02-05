package domain

import "time"

type Repair struct {
	ID          string    `json:"id"`
	EntryID     string    `json:"entry_id"`
	VehicleID   string    `json:"vehicle_id"`
	FaultID     string    `json:"fault_id,omitempty"`
	Description string    `json:"description"`
	LaborHours  float64   `json:"labor_hours"`
	PartsCost   float64   `json:"parts_cost"`
	Status      string    `json:"status"` // pending, in_progress, completed
	CompletedAt *time.Time `json:"completed_at,omitempty"`
	TechnicianID string   `json:"technician_id,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
