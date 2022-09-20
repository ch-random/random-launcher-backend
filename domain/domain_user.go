package domain

import (
	"time"
)

type User struct {
	ID        uint `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	GoogleID string `validate:"required" json:"google_id"`
	Role     string `json:"role"`
	Name     string `json:"name"`
}

type UserRepository interface {
	GetByID(id uint) (User, error)
}
