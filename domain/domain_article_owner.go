package domain

import (
	"github.com/google/uuid"
)

type ArticleOwner struct {
	ID        uuid.UUID `gorm:"type:char(36);primaryKey;not null" json:"id"`
	ArticleID uuid.UUID `gorm:"type:char(36);not null" json:"article_id"`
	// has one
	// User User `gorm:"foreignKey:ID" json:"-"`
	UserID uuid.UUID `gorm:"type:char(36);not null" json:"user_id"`
	User   User      `gorm:"foreignKey:UserID" json:"user"`
}

func (*ArticleOwner) TableName() string {
	return "article_owners"
}

type ArticleOwnerRepository interface {
	GetByID(id uuid.UUID) (ArticleOwner, error)
	GetByArticleID(id uuid.UUID) ([]ArticleOwner, error)
	Insert(ao *ArticleOwner) error
	Update(ao *ArticleOwner) error
	Delete(id uuid.UUID) error
}
