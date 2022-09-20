package pscale

import (
	"log"
	"strconv"

	"gorm.io/gorm"

	"github.com/ch-random/random-launcher-backend/domain"
	"github.com/ch-random/random-launcher-backend/repository"
)

type pscaleArticleRepository struct {
	db *gorm.DB
}

// NewPscaleArticleRepository will create an object that represent the article.Repository interface
func NewPscaleArticleRepository(db *gorm.DB) domain.ArticleRepository {
	return &pscaleArticleRepository{db}
}

func (articleRepo *pscaleArticleRepository) Fetch(cursor string, numString string) (articles []domain.Article, nextCursor string, err error) {
	// query := `SELECT id, title, content, user_id, updated_at, created_at
	// FROM article WHERE created_at > ? ORDER BY created_at LIMIT ?`
	log.Println("cursor:", cursor)
	db := articleRepo.db

	if cursor != "" {
		decodedCursor, err := repository.DecodeCursor(cursor)
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
		db = db.Limit(num);
	}

	if err = db.Find(&articles).Error; err != nil {
		return
	}

	if len(articles) == int(num) {
		nextCursor = repository.EncodeCursor(articles[len(articles)-1].CreatedAt)
	}
	return
}
func (articleRepo *pscaleArticleRepository) GetByID(id uint) (article domain.Article, err error) {
	// query := `SELECT id, title, content, user_id, updated_at, created_at
	// FROM article WHERE ID = ?`
	if result := articleRepo.db.Where("id = ?", id).First(&article); result.Error != nil {
		err = domain.ErrNotFound
	}
	return
}

func (articleRepo *pscaleArticleRepository) GetByTitle(title string) (article domain.Article, err error) {
	// query := `SELECT id, title, content, user_id, updated_at, created_at
	// FROM article WHERE title = ?`
	if result := articleRepo.db.Where("title = ?", title).First(&article); result.Error != nil {
		err = domain.ErrNotFound
	}
	return
}

func (articleRepo *pscaleArticleRepository) Insert(a *domain.Article) (err error) {
	// query := "INSERT article SET title=?, content=?, user_id=?, updated_at=?, created_at=?"
	err = articleRepo.db.Create(a).Error;
	return
}

func (articleRepo *pscaleArticleRepository) Update(a *domain.Article) (err error) {
	// query := "UPDATE article set title=?, content=?, user_id=?, updated_at=? WHERE id = ?"
	result := articleRepo.db.Updates(a);
	if err = result.Error; err != nil {
		return
	}

	affected := result.RowsAffected
	if affected != 1 {
		err = domain.ErrRowsAffectedNotOne
	}
	return
}

func (articleRepo *pscaleArticleRepository) Delete(id uint) (err error) {
	// query := "DELETE FROM article WHERE id = ?"
	result := articleRepo.db.Where("id = ?", id).Delete(&domain.Article{});
	if err = result.Error; err != nil {
		return
	}

	affected := result.RowsAffected
	if affected != 1 {
		err = domain.ErrRowsAffectedNotOne
	}
	return
}
