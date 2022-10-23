package domain

import (
	"github.com/google/uuid"
)

type ArticleImageURL struct {
	ID        uuid.UUID `gorm:"type:char(36);primaryKey;not null" json:"id"`
	ArticleID uuid.UUID `gorm:"type:char(36);not null" json:"article_id"`
	ImageURL  string    `gorm:"type:text" validate:"required" json:"image_url"`
}

type ArticleImageURLRepository interface {
	GetByID(id uuid.UUID) (ArticleImageURL, error)
	GetByArticleID(id uuid.UUID) ([]ArticleImageURL, error)
	Insert(aiu *ArticleImageURL) error
	Update(aiu *ArticleImageURL) error
	Delete(id uuid.UUID) error
}
