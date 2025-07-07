package models

import (
	"gorm.io/gorm"
	"time"
)

type Room struct {
	gorm.Model
	ID           int       `json:"room_id" gorm:"primaryKey"`
	TypeID       int       `json:"type_id" gorm:"foreignKey:TypeID"` // Foreign Key
	RoomNumber   int       `json:"room_number" validate:"required"`
	Price        float64   `json:"price" validate:"required"`
	Total        float64   `json:"total" validate:"required"`
	Availability string    `json:"availability" validate:"required"`
	Description  string    `json:"description" validate:"required"`
	Image        string    `json:"image" validate:"required"`
	CreateDate   time.Time `json:"create_date"`
	UpdateDate   time.Time `json:"update_date"`
}
