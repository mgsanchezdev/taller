package domain

import "time"

type InspectionTemplate struct {
	ID        string    `json:"id"`
	WorkshopID string   `json:"workshop_id"`
	Name      string    `json:"name"`
	Items     []InspectionItem `json:"items,omitempty"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
