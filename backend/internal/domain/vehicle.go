package domain

import "time"

type Vehicle struct {
	ID         string    `json:"id"`
	WorkshopID string    `json:"workshop_id"`
	Plate      string    `json:"plate"`
	Brand      string    `json:"brand"`
	Model      string    `json:"model"`
	Year       int       `json:"year"`
	VIN        string    `json:"vin"`
	Color      string    `json:"color"`
	OwnerName  string    `json:"owner_name"`
	OwnerPhone string    `json:"owner_phone"`
	OwnerEmail string    `json:"owner_email"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
