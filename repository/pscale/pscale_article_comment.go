package pscale

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"github.com/google/uuid"

	"github.com/ch-random/random-launcher-backend/domain"
)

type pscaleArticleCommentRepository struct {
	db *gorm.DB
}

func NewArticleCommentRepository(db *gorm.DB) domain.ArticleCommentRepository {
	return &pscaleArticleCommentRepository{db.Model(&domain.ArticleComment{})}
}

func (acr *pscaleArticleCommentRepository) GetByID(id uuid.UUID) (ac domain.ArticleComment, err error) {
	db := acr.db.Where("id = ?", id).Preload(clause.Associations).Begin().First(&ac)
	if db.Error != nil {
		err = domain.ErrNotFound
	}
	return
}

func (acr *pscaleArticleCommentRepository) GetByArticleID(id uuid.UUID) (acs []domain.ArticleComment, err error) {
	db := acr.db.Where("article_id = ?", id).Preload(clause.Associations).Begin().Find(&acs)
	if db.Error != nil {
		err = domain.ErrNotFound
	}
	return
}

func (acr *pscaleArticleCommentRepository) Insert(ac *domain.ArticleComment) (err error) {
	db := acr.db.Create(ac);
	if err = db.Error; err != nil {
		db.Rollback()
		return
	}
	return
}

func (acr *pscaleArticleCommentRepository) Update(ac *domain.ArticleComment) (err error) {
	db := acr.db.Where("id = ?", ac.ID).Updates(ac);
	if err = db.Error; err != nil {
		return
	}

	affected := db.RowsAffected
	if affected != 1 {
		err = domain.ErrRowsAffectedNotOne
	}
	return
}

func (acr *pscaleArticleCommentRepository) Delete(id uuid.UUID) (err error) {
	db := acr.db.Where("id = ?", id).Delete(&domain.ArticleComment{});
	if err = db.Error; err != nil {
		return
	}

	affected := db.RowsAffected
	if affected != 1 {
		err = domain.ErrRowsAffectedNotOne
	}
	return
}
