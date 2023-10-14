package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"type:char(36);primary_key;not null" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Role      string    `gorm:"type:text" validate:"required" json:"role"`
	Name      string    `gorm:"type:text" validate:"required" json:"name"`
}
func (*User) TableName() string {
	return "users"
}

type UserUsecase interface {
	Fetch(c context.Context, cursor string, numString string) ([]User, string, error)
	GetByID(c context.Context, id uuid.UUID) (User, error)
	Insert(c context.Context, u *User) error
	Update(c context.Context, u *User) error
}
type UserRepository interface {
	Fetch(cursor string, numString string) ([]User, string, error)
	GetByID(id uuid.UUID) (User, error)
	Insert(u *User) error
	Update(u *User) error
}
