// 【Go】Gormの使い方(CRUD)
// https://zenn.dev/a_ichi1/articles/4b113d4c46857a
package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// https://stackoverflow.com/questions/66810464/unsupported-relations-in-gorm
// https://zenn.dev/skanehira/articles/2020-09-19-go-echo-bind-tips
type Article struct {
	ID        uuid.UUID `gorm:"type:char(36);primary_key;not null" param:"id" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	EventId   string    `gorm:"type:char(36);not null" validate:"required" json:"event_id"`
	Title     string    `gorm:"type:text" validate:"required" json:"title"`
	Body      string    `gorm:"type:text" validate:"required" json:"body"`
	Public    bool      `gorm:"type:boolean" validate:"required" json:"public"`
	// belongs to
	UserID uuid.UUID `gorm:"type:char(36);not null" json:"user_id"`
	User   User      `gorm:"PRELOAD:false" json:"user"`
	// has one
	ArticleGameContent ArticleGameContent `gorm:"foreignKey:ID" json:"article_game_content"`
	// has many
	ArticleOwners    []ArticleOwner    `json:"article_owners"`
	ArticleTags      []ArticleTag      `json:"article_tags"`
	ArticleComments  []ArticleComment  `json:"article_comments"`
	ArticleImageURLs []ArticleImageURL `json:"article_image_urls"`
}
func (*Article) TableName() string {
	return "articles"
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
