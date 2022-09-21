package domain

import (
	"context"

	"github.com/google/uuid"
)

type ArticleTag struct {
	ID        uuid.UUID `gorm:"type:uuid; primaryKey" json:"id"`
	ArticleID uuid.UUID `gorm:"type:uuid" validate:"required" json:"articleId"`
	Name      string    `json:"name"`
}

type ArticleTagUsecase interface {
	GetByID(c context.Context, id uuid.UUID) (ArticleTag, error)
	GetByArticleID(c context.Context, id uuid.UUID) ([]ArticleTag, error)
	Insert(c context.Context, at *ArticleTag) error
	Update(c context.Context, at *ArticleTag) error
	Delete(c context.Context, id uuid.UUID) error
}
type ArticleTagRepository interface {
	GetByID(id uuid.UUID) (ArticleTag, error)
	GetByArticleID(id uuid.UUID) ([]ArticleTag, error)
	Insert(at *ArticleTag) error
	Update(at *ArticleTag) error
	Delete(id uuid.UUID) error
}
