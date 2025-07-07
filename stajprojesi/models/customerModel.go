package models

import (
	"gorm.io/gorm"
	"time"
)

type Customer struct {
	gorm.Model
	ID            int       `json:"customer_id" gorm:"primaryKey"`
	UserID        int       `json:"user_id" gorm:"foreignKey:UserID"` // Foreign Key
	ReservationID int       `json:"reservation_id" `
	FirstName     string    `json:"first_name"`
	LastName      string    `json:"last_name"`
	Email         string    `json:"email"`
	Phone         string    `json:"phone"`
	Address       string    `json:"address"`
	Password      string    `json:"-"`
	PasswordHash  string    `json:"-"`
	Role          string    `json:"role" `
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	Token         string    `json:"token"`
	RefreshToken  string    `json:"refresh_token"`
	Avatar        string    `json:"avatar"`
}
