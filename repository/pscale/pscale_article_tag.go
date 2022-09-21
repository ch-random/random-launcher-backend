package pscale

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"github.com/google/uuid"

	"github.com/ch-random/random-launcher-backend/domain"
)

type pscaleArticleTagRepository struct {
	db *gorm.DB
}

func NewArticleTagRepository(db *gorm.DB) domain.ArticleTagRepository {
	return &pscaleArticleTagRepository{db.Model(&domain.ArticleTag{})}
}

func (acr *pscaleArticleTagRepository) GetByID(id uuid.UUID) (at domain.ArticleTag, err error) {
	db := acr.db.Where("id = ?", id).Preload(clause.Associations).Begin().First(&at)
	if db.Error != nil {
		err = domain.ErrNotFound
	}
	return
}

func (acr *pscaleArticleTagRepository) GetByArticleID(id uuid.UUID) (ats []domain.ArticleTag, err error) {
	db := acr.db.Where("article_id = ?", id).Preload(clause.Associations).Begin().Find(&ats)
	if db.Error != nil {
		err = domain.ErrNotFound
	}
	return
}

func (acr *pscaleArticleTagRepository) Insert(at *domain.ArticleTag) (err error) {
	err = acr.db.Create(at).Error;
	return
}

func (acr *pscaleArticleTagRepository) Update(at *domain.ArticleTag) (err error) {
	db := acr.db.Updates(at);
	if err = db.Error; err != nil {
		return
	}

	affected := db.RowsAffected
	if affected != 1 {
		err = domain.ErrRowsAffectedNotOne
	}
	return
}

func (acr *pscaleArticleTagRepository) Delete(id uuid.UUID) (err error) {
	db := acr.db.Where("id = ?", id).Delete(&domain.ArticleTag{});
	if err = db.Error; err != nil {
		return
	}

	affected := db.RowsAffected
	if affected != 1 {
		err = domain.ErrRowsAffectedNotOne
	}
	return
}
