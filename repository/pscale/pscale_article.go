package pscale

import (
	"strconv"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"

	"github.com/ch-random/random-launcher-backend/domain"
	"github.com/ch-random/random-launcher-backend/repository"
)

type pscaleArticleRepository struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) domain.ArticleRepository {
	return &pscaleArticleRepository{db.Model(&domain.Article{})}
}

func (articleRepo *pscaleArticleRepository) Fetch(cursor string, numString string) (ars []domain.Article, nextCursor string, err error) {
	db := articleRepo.db

	if cursor != "" {
		decodedCursor, err := repository.DecodeCursor(cursor)
		log.Printf("decodedCursor: %v", decodedCursor)
		if err != nil && cursor != "" {
			return nil, "", domain.ErrBadParamInput
		}
		db = db.Where("created_at > ?", decodedCursor).Order("created_at")
		if err = db.Error; err != nil {
			return nil, "", err
		}
	}

	num, err := strconv.Atoi(numString)
	if err == nil {
		log.Printf("num: %v", num)
		db = db.Limit(num)
	}

	// Preloading (Eager Loading) https://gorm.io/ja_JP/docs/preload.html
	// Join Preload は、one-to-one (has one, belongs to) でのみ動作します。
	// Preloading https://qiita.com/tsubasaozawa/items/ac5a8a515fe4f7a139b0
	// Nested Preloading: db.Preload("ArticleOwner.User")
	// db.Preload(clause.Associations) により全データを preload できる。
	// 不要データを含むなら、db.Preload(""), `gorm:"PRELOAD:false"` として高速化。
	db = db.Preload(clause.Associations).Begin().Find(&ars)
	if err = db.Error; err != nil {
		return
	}
	log.Printf("ars: %v", ars)

	if len(ars) > 0 && len(ars) == num {
		nextCursor = repository.EncodeCursor(ars[len(ars)-1].CreatedAt)
	}
	return
}
func (articleRepo *pscaleArticleRepository) GetByID(id uuid.UUID) (ar domain.Article, err error) {
	db := articleRepo.db.Where("id = ?", id).Preload(clause.Associations).Begin().First(&ar)
	if db.Error != nil {
		err = domain.ErrNotFound
	}
	return
}

func (articleRepo *pscaleArticleRepository) GetByTitle(title string) (ar domain.Article, err error) {
	db := articleRepo.db.Where("title = ?", title).Preload(clause.Associations).Begin().First(&ar)
	if db.Error != nil {
		err = domain.ErrNotFound
	}
	return
}

func (articleRepo *pscaleArticleRepository) Insert(ar *domain.Article) (err error) {
	db := articleRepo.db.Create(ar)
	if err = db.Error; err != nil {
		db.Rollback()
		return
	}
	return
}

func (articleRepo *pscaleArticleRepository) Update(ar *domain.Article) (err error) {
	db := articleRepo.db.Where("id = ?", ar.ID).Updates(ar)
	if err = db.Error; err != nil {
		return
	}

	if affected := db.RowsAffected; affected != 1 {
		err = domain.ErrRowsAffectedNotOne
	}
	return
}

func (articleRepo *pscaleArticleRepository) Delete(id uuid.UUID) (err error) {
	db := articleRepo.db.Where("id = ?", id).Delete(&domain.Article{})
	if err = db.Error; err != nil {
		return
	}

	if affected := db.RowsAffected; affected != 1 {
		err = domain.ErrRowsAffectedNotOne
	}
	return
}
