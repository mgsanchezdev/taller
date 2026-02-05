package domain

import "time"

type Notification struct {
	ID        string    `json:"id"`
	EntryID   string    `json:"entry_id,omitempty"`
	VehicleID string    `json:"vehicle_id,omitempty"`
	Channel   string    `json:"channel"` // whatsapp, email, sms
	To        string    `json:"to"`
	Subject   string    `json:"subject,omitempty"`
	Body      string    `json:"body"`
	Status    string    `json:"status"` // pending, sent, failed
	SentAt    *time.Time `json:"sent_at,omitempty"`
	Error     string    `json:"error,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}
