package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"type:char(36);primary_key;not null" json:"id"`
	CreatedAt time.Time `gorm:"type:DATETIME(6);autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:DATETIME(6);autoUpdateTime" json:"updated_at"`
	GoogleID  string    `gorm:"type:char(28);not null" json:"google_id"`
	Role      string    `gorm:"type:text" validate:"required" json:"role"`
	Name      string    `gorm:"type:text" validate:"required" json:"name"`
	// Role      string    `gorm:"type:text" json:"role"`
	// Name      string    `gorm:"type:text" json:"name"`
}

type UserUsecase interface {
	Fetch(c context.Context, cursor string, numString string) ([]User, string, error)
	GetByID(c context.Context, id uuid.UUID) (User, error)
	Update(c context.Context, u *User) error
}
type UserRepository interface {
	Fetch(cursor string, numString string) ([]User, string, error)
	GetByID(id uuid.UUID) (User, error)
	Update(u *User) error
}
