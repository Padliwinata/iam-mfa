package models

import (
	uuid "github.com/satori/go.uuid"
)

type User struct {
	ID       uuid.UUID `gorm:"type:uuid;primary_key;"`
	Name     string    `gorm:"type:varchar(255);not null"`
	Email    string    `gorm:"uniqueIndex;not null"`
	Password string    `gorm:"not null"`

	Otp_enabled  bool `gorm:"default:false;"`
	Otp_verified bool `gorm:"default:false;"`

	Otp_secret   string
	Otp_auth_url string
}
type RegisterUserInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" bindinig:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginUserInput struct {
	Email    string `json:"email" bindinig:"required"`
	Password string `json:"password" binding:"required"`
}

type OTPInput struct {
	UserId string `json:"user_id"`
	Token  string `json:"token"`
}

type GetUserResponse struct {
	Id       string `json:"id,omitempty"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
