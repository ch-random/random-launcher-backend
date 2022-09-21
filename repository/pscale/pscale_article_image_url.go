package pscale

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"github.com/google/uuid"

	"github.com/ch-random/random-launcher-backend/domain"
)

type pscaleArticleImageURLRepository struct {
	db *gorm.DB
}

func NewArticleImageURLRepository(db *gorm.DB) domain.ArticleImageURLRepository {
	return &pscaleArticleImageURLRepository{db.Model(&domain.ArticleImageURL{})}
}

func (acr *pscaleArticleImageURLRepository) GetByID(id uuid.UUID) (aiu domain.ArticleImageURL, err error) {
	db := acr.db.Where("id = ?", id).Preload(clause.Associations).Begin().First(&aiu)
	if db.Error != nil {
		err = domain.ErrNotFound
	}
	return
}

func (acr *pscaleArticleImageURLRepository) GetByArticleID(id uuid.UUID) (aius []domain.ArticleImageURL, err error) {
	db := acr.db.Where("article_id = ?", id).Preload(clause.Associations).Begin().Find(&aius)
	if db.Error != nil {
		err = domain.ErrNotFound
	}
	return
}

func (acr *pscaleArticleImageURLRepository) Insert(aiu *domain.ArticleImageURL) (err error) {
	db := acr.db.Create(aiu);
	if err = db.Error; err != nil {
		db.Rollback()
		return
	}
	return
}

func (acr *pscaleArticleImageURLRepository) Update(aiu *domain.ArticleImageURL) (err error) {
	db := acr.db.Where("id = ?", aiu.ID).Updates(aiu);
	if err = db.Error; err != nil {
		return
	}

	affected := db.RowsAffected
	if affected != 1 {
		err = domain.ErrRowsAffectedNotOne
	}
	return
}

func (acr *pscaleArticleImageURLRepository) Delete(id uuid.UUID) (err error) {
	db := acr.db.Where("id = ?", id).Delete(&domain.ArticleImageURL{});
	if err = db.Error; err != nil {
		return
	}

	affected := db.RowsAffected
	if affected != 1 {
		err = domain.ErrRowsAffectedNotOne
	}
	return
}
