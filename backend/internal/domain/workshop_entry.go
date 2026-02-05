package domain

import "time"

type WorkshopEntry struct {
	ID          string    `json:"id"`
	WorkshopID  string    `json:"workshop_id"`
	VehicleID   string    `json:"vehicle_id"`
	EntryNumber string    `json:"entry_number"`
	EntryDate   time.Time `json:"entry_date"`
	Mileage     int       `json:"mileage"`
	FuelLevel   string    `json:"fuel_level"`
	Notes       string    `json:"notes"`
	Status      string    `json:"status"` // received, in_progress, completed
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
