package domain

import (
	"context"

	"github.com/google/uuid"
)

type ArticleOwner struct {
	// UserID
	ID        uuid.UUID `gorm:"type:uuid; primaryKey" json:"id"`
	ArticleID uuid.UUID `gorm:"type:uuid" validate:"required" json:"articleId"`
	// has one
	User User `gorm:"foreignKey:ID" json:"user"`
}

type ArticleOwnerUsecase interface {
	GetByID(c context.Context, id uuid.UUID) (ArticleOwner, error)
	GetByArticleID(c context.Context, id uuid.UUID) ([]ArticleOwner, error)
	Insert(c context.Context, ao *ArticleOwner) error
	Update(c context.Context, ao *ArticleOwner) error
	Delete(c context.Context, id uuid.UUID) error
}
type ArticleOwnerRepository interface {
	GetByID(id uuid.UUID) (ArticleOwner, error)
	GetByArticleID(id uuid.UUID) ([]ArticleOwner, error)
	Insert(ao *ArticleOwner) error
	Update(ao *ArticleOwner) error
	Delete(id uuid.UUID) error
}
