package usecase

import (
	"context"
	"log"
	"time"

	"golang.org/x/sync/errgroup"

	"github.com/ch-random/random-launcher-backend/domain"
)

type articleUsecase struct {
	articleRepo domain.ArticleRepository
	userRepo    domain.UserRepository
	timeout     time.Duration
}

func NewArticleUsecase(articleRepo domain.ArticleRepository, userRepo domain.UserRepository, timeout time.Duration) domain.ArticleUsecase {
	return &articleUsecase{
		articleRepo,
		userRepo,
		timeout,
	}
}

func (au *articleUsecase) fillUserDetails(c context.Context, data []domain.Article) ([]domain.Article, error) {
	eg, ctx := errgroup.WithContext(c)

	// Get the user's id
	mapUsers := map[uint]domain.User{}

	for _, article := range data {
		mapUsers[article.UserID] = domain.User{}
		log.Println("article:", article)
	}

	// Using goroutine to fetch the user's detail
	chanUser := make(chan domain.User)
	for userID := range mapUsers {
		userID := userID
		log.Println("userID:", userID)
		// 複数のGoroutineをWaitGroup（ErrGroup）で制御する
		// https://blog.toshimaru.net/goroutine-with-waitgroup/
		eg.Go(func() (err error) {
			user, err := au.userRepo.GetByID(userID)
			if err != nil {
				return
			}
			select {
			case chanUser <- user:
				return
			case <-ctx.Done():
				log.Println("the gorutine canceled. userID:", userID)
				return ctx.Err()
			}
		})
	}
	// https://pkg.go.dev/golang.org/x/sync/errgroup?utm_source=godoc#example-Group-Pipeline
	go func() {
		eg.Wait()
		close(chanUser)
	}()

	for user := range chanUser {
		if user != (domain.User{}) {
			mapUsers[user.ID] = user
		}
	}

	// merge the user's data
	for index, item := range data {
		if au, ok := mapUsers[item.User.ID]; ok {
			data[index].User = au
		}
	}
	return data, nil
}

func (au *articleUsecase) Fetch(c context.Context, cursor string, numString string) (res []domain.Article, nextCursor string, err error) {
	ctx, cancel := context.WithTimeout(c, au.timeout)
	defer cancel()

	res, nextCursor, err = au.articleRepo.Fetch(cursor, numString)
	if err != nil {
		return nil, "", err
	}

	res, err = au.fillUserDetails(ctx, res)
	if err != nil {
		nextCursor = ""
	}
	return
}

func (au *articleUsecase) GetByID(c context.Context, id uint) (res domain.Article, err error) {
	_, cancel := context.WithTimeout(c, au.timeout)
	defer cancel()

	res, err = au.articleRepo.GetByID(id)
	if err != nil {
		return
	}

	resUser, err := au.userRepo.GetByID(res.UserID)
	if err != nil {
		return domain.Article{}, err
	}
	res.User = resUser
	return
}

func (au *articleUsecase) Update(c context.Context, ar *domain.Article) error {
	_, cancel := context.WithTimeout(c, au.timeout)
	defer cancel()

	ar.UpdatedAt = time.Now()
	return au.articleRepo.Update(ar)
}

func (au *articleUsecase) GetByTitle(c context.Context, title string) (res domain.Article, err error) {
	_, cancel := context.WithTimeout(c, au.timeout)
	defer cancel()
	res, err = au.articleRepo.GetByTitle(title)
	if err != nil {
		return domain.Article{}, err
	}

	resUser, err := au.userRepo.GetByID(res.User.ID)
	if err != nil {
		return domain.Article{}, err
	}

	res.User = resUser
	return
}

func (au *articleUsecase) Insert(c context.Context, m *domain.Article) (err error) {
	ctx, cancel := context.WithTimeout(c, au.timeout)
	defer cancel()
	_, err = au.GetByTitle(ctx, m.Title)
	if err != nil {
		return err
	}
	return au.articleRepo.Insert(m)
}

func (au *articleUsecase) Delete(c context.Context, id uint) (err error) {
	_, cancel := context.WithTimeout(c, au.timeout)
	defer cancel()
	_, err = au.articleRepo.GetByID(id)
	if err != nil {
		return err
	}
	return au.articleRepo.Delete(id)
}
