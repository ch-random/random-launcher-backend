// 【Go】Gormの使い方(CRUD)
// https://zenn.dev/a_ichi1/articles/4b113d4c46857a
package domain

import (
	"context"
	"time"
)

// User belongs to Company (many-to-one)
// User has one CreditCard (one-to-one)
// User has many CreditCards (one-to-many)
// User has and belongs to many languages (many-to-many)
type Article struct {
	// https://zenn.dev/skanehira/articles/2020-09-19-go-echo-bind-tips
	// `param:"id"`: c.Param("id")
	ID        uint      `gorm:"primary_key" param:"id" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	Public    bool      `json:"public"`
	// belongs to
	UserID uint `json:"user_id"`
	User   User `json:"user"`
	// has many
	ArticleOwners    []ArticleOwner    `json:"article_owners"`
	ArticleTags      []ArticleTag      `json:"article_tags"`
	ArticleComments  []ArticleComment  `json:"article_comments"`
	ArticleImageURLs []ArticleImageURL `json:"article_image_urls"`
}
type ArticleOwner struct {
	// UserID
	ID        uint `gorm:"primaryKey" json:"id"`
	ArticleID uint `validate:"required" json:"article_id"`
	// has one
	User User `gorm:"foreignKey:ID" json:"user"`
}

// Tag 検索の実装は面倒なのでサボる
type ArticleTag struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	ArticleID uint   `validate:"required" json:"article_id"`
	Name      string `json:"name"`
}
type ArticleComment struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	ArticleID uint   `validate:"required" json:"article_id"`
	Body      string `json:"body"`
	Rate      string `json:"rate"`
}
type ArticleImageURL struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	ArticleID uint   `validate:"required" json:"article_id"`
	ImageURL  string `json:"image_url"`
}

type ArticleUsecase interface {
	Fetch(c context.Context, cursor string, numString string) ([]Article, string, error)
	GetByID(c context.Context, id uint) (Article, error)
	GetByTitle(c context.Context, title string) (Article, error)
	Insert(c context.Context, ar *Article) error
	Update(c context.Context, ar *Article) error
	Delete(c context.Context, id uint) error
}
type ArticleRepository interface {
	Fetch(cursor string, numString string) (articles []Article, nextCursor string, err error)
	GetByID(id uint) (Article, error)
	GetByTitle(title string) (Article, error)
	Insert(ar *Article) error
	Update(ar *Article) error
	Delete(id uint) error
}
