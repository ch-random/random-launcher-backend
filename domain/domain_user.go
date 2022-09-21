package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid; primary_key" json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	GoogleID string `validate:"required" json:"googleId"`
	Role     string `json:"role"`
	Name     string `json:"name"`
}

type UserRepository interface {
	GetByID(id uuid.UUID) (User, error)
}
