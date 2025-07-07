package models

import (
	"gorm.io/gorm"
	"time"
)

type RoomType struct {
	gorm.Model
	ID            int     `json:"type_id" gorm:"primaryKey"`
	TypeName      string  `json:"type_name"`
	Description   string  `json:"description"`
	NumberOfRooms int     `json:"number_of_rooms"`
	Price         float64 `json:"price"`
	CreateDate    time.Time
	UpdateDate    time.Time
}
