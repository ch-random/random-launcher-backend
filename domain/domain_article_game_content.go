package domain

import (
	"time"

	"github.com/google/uuid"
)

type ArticleGameContent struct {
	// ArticleID
	ID        uuid.UUID `gorm:"type:char(36);primaryKey;not null" validate:"required" json:"id"`
	CreatedAt time.Time `gorm:"type:DATETIME(6);autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time `gorm:"type:DATETIME(6);autoUpdateTime" json:"updatedAt"`
	ExecPath  string    `gorm:"type:text" json:"execPath"`
	ZipURL    string    `gorm:"type:text" json:"zipUrl"`
}

type ArticleGameContentRepository interface {
	GetByID(id uuid.UUID) (ArticleGameContent, error)
	GetByArticleID(id uuid.UUID) (ArticleGameContent, error)
	Insert(agc *ArticleGameContent) error
	Update(agc *ArticleGameContent) error
	Delete(id uuid.UUID) error
}
