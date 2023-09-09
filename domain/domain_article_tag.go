package domain

import (
	"github.com/google/uuid"
)

type ArticleTag struct {
	ID        uuid.UUID `gorm:"type:char(36);primaryKey;not null" json:"id"`
	ArticleID uuid.UUID `gorm:"type:char(36);not null" json:"article_id"`
	Name      string    `gorm:"type:text" validate:"required" json:"name"`
}
func (*ArticleTag) TableName() string {
	return "article_tags"
}

type ArticleTagRepository interface {
	GetByID(id uuid.UUID) (ArticleTag, error)
	GetByArticleID(id uuid.UUID) ([]ArticleTag, error)
	Insert(at *ArticleTag) error
	Update(at *ArticleTag) error
	Delete(id uuid.UUID) error
}
