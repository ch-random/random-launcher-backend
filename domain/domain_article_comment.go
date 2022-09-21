package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type ArticleComment struct {
	ID        uuid.UUID   `gorm:"type:uuid; primaryKey" json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	ArticleID uuid.UUID   `gorm:"type:uuid" validate:"required" json:"articleId"`
	Body      string `json:"body"`
	Rate      string `json:"rate"`
}

type ArticleCommentUsecase interface {
	GetByID(c context.Context, id uuid.UUID) (ArticleComment, error)
	GetByArticleID(c context.Context, id uuid.UUID) ([]ArticleComment, error)
	Insert(c context.Context, ac *ArticleComment) error
	Update(c context.Context, ac *ArticleComment) error
	Delete(c context.Context, id uuid.UUID) error
}
type ArticleCommentRepository interface {
	GetByID(id uuid.UUID) (ArticleComment, error)
	GetByArticleID(id uuid.UUID) ([]ArticleComment, error)
	Insert(ac *ArticleComment) error
	Update(ac *ArticleComment) error
	Delete(id uuid.UUID) error
}
