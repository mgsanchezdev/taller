package domain

import "time"

type InspectionItem struct {
	ID         string    `json:"id"`
	TemplateID string    `json:"template_id"`
	Label      string    `json:"label"`
	Category   string    `json:"category"`
	Order      int       `json:"order"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type InspectionItemResult struct {
	ID           string `json:"id"`
	InspectionID string `json:"inspection_id"`
	ItemID       string `json:"item_id"`
	Status       string `json:"status"` // ok, damaged, missing, n/a
	Notes        string `json:"notes"`
	PhotoURL     string `json:"photo_url,omitempty"`
}
