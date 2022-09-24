package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"type:char(36);primary_key;not null" json:"id"`
	CreatedAt time.Time `gorm:"type:DATETIME(6);autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time `gorm:"type:DATETIME(6);autoUpdateTime" json:"updatedAt"`
	GoogleID  string    `gorm:"type:char(28);not null" json:"googleId"`
	Role      string    `gorm:"type:text" validate:"required" json:"role"`
	Name      string    `gorm:"type:text" validate:"required" json:"name"`
}

type UserRepository interface {
	GetByID(id uuid.UUID) (User, error)
}
