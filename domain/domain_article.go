// 【Go】Gormの使い方(CRUD)
// https://zenn.dev/a_ichi1/articles/4b113d4c46857a
package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// User belongs to Company (many-to-one)
// User has one CreditCard (one-to-one)
// User has many CreditCards (one-to-many)
// User has and belongs to many languages (many-to-many)
type Article struct {
	// https://stackoverflow.com/questions/66810464/unsupported-relations-in-gorm
	// https://zenn.dev/skanehira/articles/2020-09-19-go-echo-bind-tips
	// `param:"id"`: c.Param("id")
	ID        uuid.UUID `gorm:"type:char(36);primary_key;not null" validate:"required" param:"id" json:"Id"`
	CreatedAt time.Time `gorm:"type:DATETIME(6);autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time `gorm:"type:DATETIME(6);autoUpdateTime" json:"updatedAt"`
	Title     string    `gorm:"type:text" json:"title"`
	Body      string    `gorm:"type:text" json:"body"`
	Public    bool      `gorm:"type:boolean" json:"public"`
	// belongs to
	UserID uuid.UUID `gorm:"type:char(36);not null" validate:"required" json:"userId"`
	User   User      `json:"user"`
	// has one
	ArticleGameContent ArticleGameContent `gorm:"foreignKey:ID" json:"articleGameContents"`
	// has many
	ArticleOwners    []ArticleOwner    `json:"articleOwners,omitempty"`
	ArticleTags      []ArticleTag      `json:"articleTags,omitempty"`
	ArticleComments  []ArticleComment  `json:"articleComments,omitempty"`
	ArticleImageURLs []ArticleImageURL `json:"articleImageUrls,omitempty"`
}

type ArticleUsecase interface {
	Fetch(c context.Context, cursor string, numString string) ([]Article, string, error)
	GetByID(c context.Context, id uuid.UUID) (Article, error)
	GetByTitle(c context.Context, title string) (Article, error)
	Insert(c context.Context, ar *Article) error
	Update(c context.Context, ar *Article) error
	Delete(c context.Context, id uuid.UUID) error
}
type ArticleRepository interface {
	Fetch(cursor string, numString string) (articles []Article, nextCursor string, err error)
	GetByID(id uuid.UUID) (Article, error)
	GetByTitle(title string) (Article, error)
	Insert(ar *Article) error
	Update(ar *Article) error
	Delete(id uuid.UUID) error
}
