package domain

import "time"

type Company struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	TaxID     string    `json:"tax_id"`
	Address   string    `json:"address"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
