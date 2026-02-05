package domain

import "time"

type Workshop struct {
	ID        string    `json:"id"`
	CompanyID string    `json:"company_id"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	Phone     string    `json:"phone"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
