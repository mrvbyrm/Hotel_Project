package models

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	//AdminID      uint     `json:"admin_id" gorm:"primaryKey"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	PasswordHash string `json:"-"`
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
	Avatar       string `json:"avatar"`
	Role         string `json:"role"`
}
