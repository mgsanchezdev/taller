package domain

import "time"

type Role string

const (
	RoleAdmin   Role = "admin"
	RoleManager Role = "manager"
	RoleTech    Role = "tech"
	RoleReception Role = "reception"
)

type User struct {
	ID         string    `json:"id"`
	WorkshopID string    `json:"workshop_id"`
	Email      string    `json:"email"`
	PasswordHash string  `json:"-"`
	Name       string    `json:"name"`
	Role       Role      `json:"role"`
	Active     bool      `json:"active"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
