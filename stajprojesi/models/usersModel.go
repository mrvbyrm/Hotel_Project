package models

import (
	"time"
)

// User represents a user in the system
type User struct {
	ID           int       `json:"user_id" gorm:"primaryKey"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	PasswordHash string    `json:"*"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Token        string    `json:"token"`
	RefreshToken string    `json:"refresh_token"`
	Avatar       string    `json:"avatar"`
	Role         string    `json:"role"`
}
