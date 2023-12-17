package schemas

import (
	"gorm.io/gorm"
	"time"
)

type UserRegister struct {
	gorm.Model
	Name         string
	Email        string
	PasswordHash string
}

type UserRegisterResponse struct {
	ID           uint      `json:"id"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	DeletedAt    time.Time `json:"deletedAt,omitempty"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"password_hash"`
}
