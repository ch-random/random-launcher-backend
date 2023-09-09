package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type ArticleComment struct {
	ID        uuid.UUID `gorm:"type:char(36);primaryKey;not null" json:"id"`
	CreatedAt time.Time `gorm:"type:DATETIME(6);autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:DATETIME(6);autoUpdateTime" json:"updated_at"`
	ArticleID uuid.UUID `gorm:"type:char(36);not null" json:"article_id"`
	Body      string    `gorm:"type:text" validate:"required" json:"body"`
	Rate      int       `validate:"required,gte=1,lte=5" json:"rate"` // 1-5
}
func (*ArticleComment) TableName() string {
	return "article_comments"
}

type ArticleCommentUsecase interface {
	GetByID(c context.Context, id uuid.UUID) (ArticleComment, error)
	GetByArticleID(c context.Context, id uuid.UUID) ([]ArticleComment, error)
	Insert(c context.Context, ac *ArticleComment) error
	Update(c context.Context, ac *ArticleComment) error
	Delete(c context.Context, id uuid.UUID) error
	DeleteByArticleID(c context.Context, id uuid.UUID) error
}
type ArticleCommentRepository interface {
	GetByID(id uuid.UUID) (ArticleComment, error)
	GetByArticleID(id uuid.UUID) ([]ArticleComment, error)
	Insert(ac *ArticleComment) error
	Update(ac *ArticleComment) error
	Delete(id uuid.UUID) error
	DeleteByArticleID(id uuid.UUID) error
}
