package pscale

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"github.com/google/uuid"

	"github.com/ch-random/random-launcher-backend/domain"
)

type pscaleArticleGameContentRepository struct {
	db *gorm.DB
}

func NewArticleGameContentRepository(db *gorm.DB) domain.ArticleGameContentRepository {
	return &pscaleArticleGameContentRepository{db.Model(&domain.ArticleGameContent{})}
}

func (acr *pscaleArticleGameContentRepository) GetByID(id uuid.UUID) (agc domain.ArticleGameContent, err error) {
	db := acr.db.Where("id = ?", id).Preload(clause.Associations).Begin().First(&agc)
	if db.Error != nil {
		err = domain.ErrNotFound
	}
	return
}

func (acr *pscaleArticleGameContentRepository) GetByArticleID(id uuid.UUID) (agc domain.ArticleGameContent, err error) {
	db := acr.db.Where("article_id = ?", id).Preload(clause.Associations).Begin().First(&agc)
	if db.Error != nil {
		err = domain.ErrNotFound
	}
	return
}

func (acr *pscaleArticleGameContentRepository) Insert(agc *domain.ArticleGameContent) (err error) {
	db := acr.db.Create(agc);
	if err = db.Error; err != nil {
		db.Rollback()
		return
	}
	return
}

func (acr *pscaleArticleGameContentRepository) Update(agc *domain.ArticleGameContent) (err error) {
	db := acr.db.Where("id = ?", agc.ID).Updates(agc);
	if err = db.Error; err != nil {
		return
	}

	affected := db.RowsAffected
	if affected != 1 {
		err = domain.ErrRowsAffectedNotOne
	}
	return
}

func (acr *pscaleArticleGameContentRepository) Delete(id uuid.UUID) (err error) {
	db := acr.db.Where("id = ?", id).Delete(&domain.ArticleGameContent{});
	if err = db.Error; err != nil {
		return
	}

	affected := db.RowsAffected
	if affected != 1 {
		err = domain.ErrRowsAffectedNotOne
	}
	return
}
