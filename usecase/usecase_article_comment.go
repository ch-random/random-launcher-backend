package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/ch-random/random-launcher-backend/domain"
)

type articleCommentUsecase struct {
	acr     domain.ArticleCommentRepository
	timeout time.Duration
}

func NewArticleCommentUsecase(
	acr domain.ArticleCommentRepository,
	timeout time.Duration,
) domain.ArticleCommentUsecase {
	return &articleCommentUsecase{
		acr,
		timeout,
	}
}

func (acu *articleCommentUsecase) GetByID(c context.Context, id uuid.UUID) (ac domain.ArticleComment, err error) {
	_, cancel := context.WithTimeout(c, acu.timeout)
	defer cancel()

	ac, err = acu.acr.GetByID(id)
	if err != nil {
		return domain.ArticleComment{}, err
	}
	return
}

func (acu *articleCommentUsecase) GetByArticleID(c context.Context, id uuid.UUID) (acs []domain.ArticleComment, err error) {
	_, cancel := context.WithTimeout(c, acu.timeout)
	defer cancel()

	acs, err = acu.acr.GetByArticleID(id)
	if err != nil {
		return []domain.ArticleComment{}, err
	}
	return
}

func (acu *articleCommentUsecase) Update(c context.Context, ac *domain.ArticleComment) error {
	_, cancel := context.WithTimeout(c, acu.timeout)
	defer cancel()

	ac.UpdatedAt = time.Now()
	return acu.acr.Update(ac)
}

func (acu *articleCommentUsecase) Insert(c context.Context, ac *domain.ArticleComment) (err error) {
	_, cancel := context.WithTimeout(c, acu.timeout)
	defer cancel()

	if err = acu.acr.Insert(ac); err != nil {
		return err
	}
	return acu.acr.Insert(ac)
}

func (acu *articleCommentUsecase) Delete(c context.Context, id uuid.UUID) (err error) {
	_, cancel := context.WithTimeout(c, acu.timeout)
	defer cancel()

	return acu.acr.Delete(id)
}

func (acu *articleCommentUsecase) DeleteByArticleID(c context.Context, id uuid.UUID) (err error) {
	_, cancel := context.WithTimeout(c, acu.timeout)
	defer cancel()

	return acu.acr.DeleteByArticleID(id)
}
