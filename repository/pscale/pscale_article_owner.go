package pscale

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"github.com/google/uuid"

	"github.com/ch-random/random-launcher-backend/domain"
)

type pscaleArticleOwnerRepository struct {
	db *gorm.DB
}

func NewArticleOwnerRepository(db *gorm.DB) domain.ArticleOwnerRepository {
	return &pscaleArticleOwnerRepository{db.Model(&domain.ArticleOwner{})}
}

func (acr *pscaleArticleOwnerRepository) GetByID(id uuid.UUID) (ao domain.ArticleOwner, err error) {
	db := acr.db.Where("id = ?", id).First(&ao)
	if db.Error != nil {
		err = domain.ErrNotFound
	}
	return
}

func (acr *pscaleArticleOwnerRepository) GetByArticleID(id uuid.UUID) (aos []domain.ArticleOwner, err error) {
	db := acr.db.Where("article_id = ?", id).Preload(clause.Associations).Begin().Find(&aos)
	if db.Error != nil {
		err = domain.ErrNotFound
	}
	return
}

func (acr *pscaleArticleOwnerRepository) Insert(ao *domain.ArticleOwner) (err error) {
	db := acr.db.Create(ao)
	if err = db.Error; err != nil {
		db.Rollback()
		return
	}
	return
}

func (acr *pscaleArticleOwnerRepository) Update(ao *domain.ArticleOwner) (err error) {
	db := acr.db.Where("id = ?", ao.ID).Updates(ao);
	if err = db.Error; err != nil {
		return
	}

	affected := db.RowsAffected
	if affected != 1 {
		err = domain.ErrRowsAffectedNotOne
	}
	return
}

func (acr *pscaleArticleOwnerRepository) Delete(id uuid.UUID) (err error) {
	db := acr.db.Where("id = ?", id).Delete(&domain.ArticleOwner{});
	if err = db.Error; err != nil {
		return
	}

	affected := db.RowsAffected
	if affected != 1 {
		err = domain.ErrRowsAffectedNotOne
	}
	return
}
