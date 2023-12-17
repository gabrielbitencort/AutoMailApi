package schemas

import (
	"gorm.io/gorm"
	"time"
)

type UserRegister struct {
	gorm.Model
	UserId       string
	Name         string
	Email        string
	PasswordHash string
}

type UserRegisterResponse struct {
	ID           uint      `json:"id"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	DeletedAt    time.Time `json:"deletedAt,omitempty"`
	UserId       string    `json:"user_id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"password_hash"`
}
