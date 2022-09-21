package domain

import (
	"context"

	"github.com/google/uuid"
)

type ArticleImageURL struct {
	ID        uuid.UUID   `gorm:"type:uuid; primaryKey" json:"id"`
	ArticleID uuid.UUID   `gorm:"type:uuid" validate:"required" json:"articleId"`
	ImageURL  string `json:"imageUrl"`
}

type ArticleImageURLUsecase interface {
	GetByID(c context.Context, id uuid.UUID) (ArticleImageURL, error)
	GetByArticleID(c context.Context, id uuid.UUID) ([]ArticleImageURL, error)
	Insert(c context.Context, aiu *ArticleImageURL) error
	Update(c context.Context, aiu *ArticleImageURL) error
	Delete(c context.Context, id uuid.UUID) error
}
type ArticleImageURLRepository interface {
	GetByID(id uuid.UUID) (ArticleImageURL, error)
	GetByArticleID(id uuid.UUID) ([]ArticleImageURL, error)
	Insert(aiu *ArticleImageURL) error
	Update(aiu *ArticleImageURL) error
	Delete(id uuid.UUID) error
}
