package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type ArticleGameContent struct {
	// ArticleID
	ID        uuid.UUID      `gorm:"type:uuid; primaryKey" json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	ExecPath  string    `json:"execPath"`
	ZipURL    string    `json:"zipUrl"`
}

type ArticleGameContentUsecase interface {
	GetByID(c context.Context, id uuid.UUID) (ArticleGameContent, error)
	GetByArticleID(c context.Context, id uuid.UUID) (ArticleGameContent, error)
	Insert(c context.Context, agc *ArticleGameContent) error
	Update(c context.Context, agc *ArticleGameContent) error
	Delete(c context.Context, id uuid.UUID) error
}
type ArticleGameContentRepository interface {
	GetByID(id uuid.UUID) (ArticleGameContent, error)
	GetByArticleID(id uuid.UUID) (ArticleGameContent, error)
	Insert(agc *ArticleGameContent) error
	Update(agc *ArticleGameContent) error
	Delete(id uuid.UUID) error
}
