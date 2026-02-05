package domain

import "time"

type Fault struct {
	ID          string    `json:"id"`
	EntryID     string    `json:"entry_id"`
	InspectionID string   `json:"inspection_id,omitempty"`
	Description string    `json:"description"`
	Severity    string    `json:"severity"` // low, medium, high, critical
	Status      string    `json:"status"`   // open, in_repair, resolved
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
