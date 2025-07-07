package models

import (
	"gorm.io/gorm"
	"time"
)

type Reservation struct {
	gorm.Model
	ReservationID int        `json:"id"`
	UserID        int        `json:"user_id"`
	CustomerID    int        `json:"customer_id"`
	RoomID        int        `json:"room_id"`
	RoomTypeID    int        `json:"type_id"`
	CheckInDate   time.Time  `json:"check_in_date"`
	CheckOutDate  time.Time  `json:"check_out_date"`
	Status        string     `json:"status"`
	PaymentMethod string     `json:"payment_method"`
	TotalPrice    float64    `json:"total_price"`
	CreateDate    time.Time  `json:"create_date"`
	UpdateDate    *time.Time `json:"update_date,omitempty"`
	PaymentDate   *time.Time `json:"payment_date,omitempty"`
}

func GetReservationsByUserID(db *gorm.DB, userID int) ([]Reservation, error) {
	var reservations []Reservation
	if err := db.Where("customer_id = ?", userID).Find(&reservations).Error; err != nil {
		return nil, err
	}
	return reservations, nil
}

func GetReservationByID(db *gorm.DB, reservationID string) (Reservation, error) {
	var reservation Reservation
	if err := db.Where("id = ?", reservationID).First(&reservation).Error; err != nil {
		return reservation, err
	}
	return reservation, nil
}

func CreateReservation(db *gorm.DB, reservation Reservation) error {
	if err := db.Create(&reservation).Error; err != nil {
		return err
	}
	return nil
}

func UpdateReservation(db *gorm.DB, reservationID string, updatedReservation Reservation) error {
	if err := db.Model(&Reservation{}).Where("reservation_id = ?", reservationID).Updates(updatedReservation).Error; err != nil {
		return err
	}
	return nil
}

func DeleteReservation(db *gorm.DB, reservationID string) error {
	if err := db.Where("id = ?", reservationID).Delete(&Reservation{}).Error; err != nil {
		return err
	}
	return nil
}
