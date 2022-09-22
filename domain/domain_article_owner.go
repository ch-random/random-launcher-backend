package domain

import (
	"github.com/google/uuid"
)

type ArticleOwner struct {
	// UserID
	ID        uuid.UUID `gorm:"type:char(36);primaryKey;not null" validate:"required" json:"id"`
	ArticleID uuid.UUID `gorm:"type:char(36);not null" validate:"required" json:"articleId"`
	// has one
	User User `gorm:"foreignKey:ID" json:"user"`
}

type ArticleOwnerRepository interface {
	GetByID(id uuid.UUID) (ArticleOwner, error)
	GetByArticleID(id uuid.UUID) ([]ArticleOwner, error)
	Insert(ao *ArticleOwner) error
	Update(ao *ArticleOwner) error
	Delete(id uuid.UUID) error
}
