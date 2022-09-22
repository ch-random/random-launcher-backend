package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type ArticleComment struct {
	ID        uuid.UUID `gorm:"type:char(36);primaryKey;not null" validate:"required" json:"id"`
	CreatedAt time.Time `gorm:"type:DATETIME(6);autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time `gorm:"type:DATETIME(6);autoUpdateTime" json:"updatedAt"`
	ArticleID uuid.UUID `gorm:"type:char(36);not null" validate:"required" json:"articleId"`
	Body      string    `gorm:"type:text" json:"body"`
	Rate      int       `json:"rate"`
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
