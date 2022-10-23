package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"

	"github.com/ch-random/random-launcher-backend/domain"
)

type articleUsecase struct {
	ur      domain.UserRepository
	ar      domain.ArticleRepository
	agcr    domain.ArticleGameContentRepository
	aor     domain.ArticleOwnerRepository
	atr     domain.ArticleTagRepository
	acr     domain.ArticleCommentRepository
	aiur    domain.ArticleImageURLRepository
	timeout time.Duration
}

func NewArticleUsecase(
	ur domain.UserRepository,
	ar domain.ArticleRepository,
	agcr domain.ArticleGameContentRepository,
	aor domain.ArticleOwnerRepository,
	atr domain.ArticleTagRepository,
	acr domain.ArticleCommentRepository,
	aiur domain.ArticleImageURLRepository,
	timeout time.Duration,
) domain.ArticleUsecase {
	return &articleUsecase{
		ur,
		ar,
		agcr,
		aor,
		atr,
		acr,
		aiur,
		timeout,
	}
}

func (au *articleUsecase) fillArticlesDetail(c context.Context, ars []domain.Article) ([]domain.Article, error) {
	ctx, cancel := context.WithTimeout(c, au.timeout)
	defer cancel()
	eg, _ := errgroup.WithContext(ctx)

	// ars[i].ArticleOwner.User
	// aosList <- ars
	aosList := map[uuid.UUID][]domain.ArticleOwner{}
	for _, ar := range ars {
		aosList[ar.ID] = []domain.ArticleOwner{}
	}
	chArticleOwners := make(chan []domain.ArticleOwner)

	// chArticleOwners <- aosList
	eg.Go(func() (err error) {
		defer close(chArticleOwners)
		for articleID := range aosList {
			// articleID := articleID
			aos, err := au.aor.GetByArticleID(articleID)
			if err != nil {
				return err
			}
			for i, ao := range aos {
				aos[i].User, err = au.ur.GetByID(ao.ID)
				if err != nil {
					return err
				}
			}
			chArticleOwners <- aos
		}
		return
	})
	go func() {
		// https://pkg.go.dev/golang.org/x/sync/errgroup?utm_source=godoc#example-Group-Pipeline
		if err := eg.Wait(); err != nil {
			log.Err(err)
		}
	}()

	// aosList <- chArticleOwners
	for aos := range chArticleOwners {
		if len(aos) > 0 {
			aosList[aos[0].ArticleID] = aos
		}
	}

	// ars[i].ArticleOwners <- aosList
	for i, ar := range ars {
		aos, ok := aosList[ar.ID]
		ars[i].ArticleOwners = aos
		if !ok {
			log.Warn().Msgf("aosList[%d] is invalid", ar.ID)
		}
	}

	// Check whether any of the goroutines failed
	if err := eg.Wait(); err != nil {
		return []domain.Article{}, nil
	}
	return ars, nil
}
func (au *articleUsecase) Fetch(c context.Context, cursor string, numString string) (ars []domain.Article, nextCursor string, err error) {
	ctx, cancel := context.WithTimeout(c, au.timeout)
	defer cancel()

	ars, nextCursor, err = au.ar.Fetch(cursor, numString)
	if err != nil {
		return nil, "", err
	}

	ars, err = au.fillArticlesDetail(ctx, ars)
	if err != nil {
		nextCursor = ""
	}
	return
}

func (au *articleUsecase) fillArticleDetail(c context.Context, ar domain.Article) (domain.Article, error) {
	ctx, cancel := context.WithTimeout(c, au.timeout)
	defer cancel()
	eg, _ := errgroup.WithContext(ctx)

	chArticleOwners := make(chan []domain.ArticleOwner)

	eg.Go(func() (err error) {
		defer close(chArticleOwners)
		aos := ar.ArticleOwners
		for _, ao := range aos {
			ao.User, err = au.ur.GetByID(ao.ID)
			if err != nil {
				return
			}
		}
		chArticleOwners <- aos
		return
	})
	go func() {
		// https://pkg.go.dev/golang.org/x/sync/errgroup?utm_source=godoc#example-Group-Pipeline
		if err := eg.Wait(); err != nil {
			log.Err(err)
		}
	}()

	ar.ArticleOwners = <-chArticleOwners

	// Check whether any of the goroutines failed
	if err := eg.Wait(); err != nil {
		return domain.Article{}, nil
	}
	return ar, nil
}

func fillArticleIDs(ar *domain.Article) (*domain.Article, error) {
	aid := uuid.New()
	log.Printf("aid: %v", aid)
	ar.ID = aid

	uid := uuid.New()
	log.Printf("uid: %v", uid)
	ar.UserID = uid
	ar.User.ID = uid

	ar.ArticleGameContent.ID = aid

	// uuid.Nil: 00000000-0000-0000-0000-000000000000
	for i, ao := range ar.ArticleOwners {
		if ao.ID == uuid.Nil {
			aoid := uuid.New()
			log.Printf("i, aoid: %v, %v", i, aoid)
			ar.ArticleOwners[i].ID = aoid
		}
		ar.ArticleOwners[i].ArticleID = aid
	}
	for i, at := range ar.ArticleTags {
		if at.ID == uuid.Nil {
			atid := uuid.New()
			log.Printf("i, atid: %v, %v", i, atid)
			ar.ArticleTags[i].ID = atid
		}
		ar.ArticleTags[i].ArticleID = aid
	}
	for i, ac := range ar.ArticleComments {
		if ac.ID == uuid.Nil {
			acid := uuid.New()
			log.Printf("i, acid: %v, %v", i, acid)
			ar.ArticleComments[i].ID = acid
		}
		ar.ArticleComments[i].ArticleID = aid
	}
	for i, aiu := range ar.ArticleImageURLs {
		if aiu.ID == uuid.Nil {
			aiuid := uuid.New()
			log.Printf("i, aiuid: %v, %v", i, aiuid)
			ar.ArticleImageURLs[i].ID = aiuid
		}
		ar.ArticleImageURLs[i].ArticleID = aid
	}
	return ar, nil
}

func (au *articleUsecase) GetByID(c context.Context, id uuid.UUID) (ar domain.Article, err error) {
	ctx, cancel := context.WithTimeout(c, au.timeout)
	defer cancel()

	ar, err = au.ar.GetByID(id)
	if err != nil {
		return domain.Article{}, err
	}

	ar, err = au.fillArticleDetail(ctx, ar)
	return
}

func (au *articleUsecase) Update(c context.Context, ar *domain.Article) error {
	_, cancel := context.WithTimeout(c, au.timeout)
	defer cancel()

	ar.UpdatedAt = time.Now()
	return au.ar.Update(ar)
}

func (au *articleUsecase) GetByTitle(c context.Context, title string) (ar domain.Article, err error) {
	ctx, cancel := context.WithTimeout(c, au.timeout)
	defer cancel()

	ar, err = au.ar.GetByTitle(title)
	if err != nil {
		return domain.Article{}, err
	}

	ar, err = au.fillArticleDetail(ctx, ar)
	return
}

func (au *articleUsecase) Insert(c context.Context, ar *domain.Article) (err error) {
	ctx, cancel := context.WithTimeout(c, au.timeout)
	defer cancel()

	// Check for title conflicts
	_, err = au.GetByTitle(ctx, ar.Title)
	if err == nil {
		return domain.ErrConflict
	} else if err != domain.ErrNotFound {
		return err
	}
	ar, err = fillArticleIDs(ar)
	if err != nil {
		return err
	}
	// au.ur.Update()
	log.Printf("ar: %v", ar)
	return au.ar.Insert(ar)
}

func (au *articleUsecase) Delete(c context.Context, id uuid.UUID) (err error) {
	_, cancel := context.WithTimeout(c, au.timeout)
	defer cancel()

	return au.ar.Delete(id)
}
