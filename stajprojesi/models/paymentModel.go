package models

import (
	"errors"
	"gorm.io/gorm"
	"time"
)

type Payment struct {
	gorm.Model
	ID            uint      `json:"id" gorm:"primaryKey"`
	Amount        float64   `json:"amount"`
	Status        string    `json:"status"`
	CreatedAt     time.Time `json:"created_at"`
	ReservationID int       `json:"reservation_id" gorm:"foreignKey:ReservationID"` // Foreign Key
	PaymentDate   time.Time `json:"payment_date"`
	PaymentMethod string    `json:"payment_method"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// GetPayment retrieves all payments from the database
func GetPayment(db *gorm.DB) ([]Payment, error) {
	var payments []Payment
	if err := db.Find(&payments).Error; err != nil {
		return nil, err
	}
	return payments, nil
}

// GetPaymentByID retrieves a single payment by its ID
func GetPaymentByID(db *gorm.DB, id uint) (Payment, error) {
	var payment Payment
	if err := db.Where("id = ?", id).First(&payment).Error; err != nil {
		return payment, err
	}
	return payment, nil
}

// CreatePayment creates a new payment in the database
func CreatePayment(db *gorm.DB, payment Payment) error {
	if err := db.Create(&payment).Error; err != nil {
		return err
	}
	return nil
}

// RefundPayment processes a refund for a specific payment
func RefundPayment(db *gorm.DB, paymentID uint) error {
	var payment Payment
	if err := db.Where("payment_id = ?", paymentID).First(&payment).Error; err != nil {
		return err
	}
	// Eğer ödeme bulunamazsa veya zaten geri iade edildiyse hata döner
	if payment.Status == "refunded" {
		return errors.New("payment already refunded")
	}
	// Ödeme durumunu güncelle
	payment.Status = "refunded"
	if err := db.Save(&payment).Error; err != nil {
		return err
	}
	return nil
}
