package version

import (
	"gorm.io/gorm"
	"time"

	"github.com/ch-random/random-launcher-backend/domain"
	"github.com/google/uuid"
)

type v20230830ArticleNewColumn struct {
	EventId string `gorm:"type:text" validate:"required" json:"event_id"`
}
func (*v20230830ArticleNewColumn) TableName() string {
	return "articles"
}

type v20230830ArticleComment struct {
	ID        uuid.UUID `gorm:"type:char(36);primaryKey;not null" json:"id"`
	CreatedAt time.Time `gorm:"type:char(6);autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:char(6);autoUpdateTime" json:"updated_at"`
	ArticleID uuid.UUID `gorm:"type:char(36);not null" json:"article_id"`
	Body      string    `gorm:"type:text" validate:"required" json:"body"`
	Rate      int       `validate:"required,gte=1,lte=5" json:"rate"` // 1-5
}

func (*v20230830ArticleComment) TableName() string {
	return "article_comments"
}
type v20230830ArticleGameContent struct {
	// ArticleID
	ID        uuid.UUID `gorm:"type:char(36);primaryKey;not null" json:"id"`
	CreatedAt time.Time `gorm:"type:char(6);autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:char(6);autoUpdateTime" json:"updated_at"`
	ExecPath  string    `gorm:"type:text" validate:"required" json:"exec_path"`
	ZipURL    string    `gorm:"type:text" validate:"required" json:"zip_url"`
}

func (*v20230830ArticleGameContent) TableName() string {
	return "article_game_contents"
}
type v20230830ArticleImageURL struct {
	ID        uuid.UUID `gorm:"type:char(36);primaryKey;not null" json:"id"`
	ArticleID uuid.UUID `gorm:"type:char(36);not null" json:"article_id"`
	ImageURL  string    `gorm:"type:text" validate:"required" json:"image_url"`
}
func (*v20230830ArticleImageURL) TableName() string {
	return "article_image_urls"
}

type v20230830ArticleOwner struct {
	// UserID
	ID        uuid.UUID `gorm:"type:char(36);primaryKey;not null" json:"id"`
	ArticleID uuid.UUID `gorm:"type:char(36);not null" json:"article_id"`
	// has one
	// User User `gorm:"foreignKey:ID" json:"-"`
	User v20230830User `gorm:"foreignKey:ID" json:"user"`
}
func (*v20230830ArticleOwner) TableName() string {
	return "article_owners"
}

type v20230830ArticleTag struct {
	ID        uuid.UUID `gorm:"type:char(36);primaryKey;not null" json:"id"`
	ArticleID uuid.UUID `gorm:"type:char(36);not null" json:"article_id"`
	Name      string    `gorm:"type:text" validate:"required" json:"name"`
}
func (*v20230830ArticleTag) TableName() string {
	return "article_tags"
}

type v20230830Article struct {
	ID        uuid.UUID `gorm:"type:char(36);primary_key;not null" param:"id" json:"id"`
	CreatedAt time.Time `gorm:"type:char(6);autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:char(6);autoUpdateTime" json:"updated_at"`
	EventId   string    `gorm:"type:char(36);not null" validate:"required" json:"event_id"`
	Title     string    `gorm:"type:text" validate:"required" json:"title"`
	Body      string    `gorm:"type:text" validate:"required" json:"body"`
	Public    bool      `gorm:"type:boolean" validate:"required" json:"public"`
	// belongs to
	UserID uuid.UUID     `gorm:"type:char(36);not null" json:"user_id"`
	User   v20230830User `gorm:"PRELOAD:false" json:"user"`
	// has one
	ArticleGameContent v20230830ArticleGameContent `gorm:"foreignKey:ID" json:"article_game_content"`
	// has many
	ArticleOwners    []v20230830ArticleOwner    `json:"article_owners"`
	ArticleTags      []v20230830ArticleTag      `json:"article_tags"`
	ArticleComments  []v20230830ArticleComment  `json:"article_comments"`
	ArticleImageURLs []v20230830ArticleImageURL `json:"article_image_urls"`
}
func (*v20230830Article) TableName() string {
	return "articles"
}

type v20230830User struct {
	ID        uuid.UUID `gorm:"type:char(36);primary_key;not null" json:"id"`
	CreatedAt time.Time `gorm:"type:char(6);autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:char(6);autoUpdateTime" json:"updated_at"`
	GoogleID  string    `gorm:"type:char(28);not null" json:"google_id"`
	Role      string    `gorm:"type:text" validate:"required" json:"role"`
	Name      string    `gorm:"type:text" validate:"required" json:"name"`
}
func (*v20230830User) TableName() string {
	return "users"
}

// 2023y: 2023 友好歳
// 2023s: 2023 白鷺歳
func V20230830() domain.Migration {
	migrate := func(db *gorm.DB) (err error) {
		if err = db.Migrator().AddColumn(&v20230830ArticleNewColumn{}, "EventId"); err != nil {
			return err
		}
		db.Model(&v20230830Article{}).Update("EventId", "2023s")
		return
	}
	rollback := func(db *gorm.DB) (err error) {
		if err = db.Migrator().DropColumn(&v20230830ArticleNewColumn{}, "EventId"); err != nil {
			return err
		}
		return
	}
	return domain.Migration{
		Migrate:  migrate,
		Rollback: rollback,
	}
}
