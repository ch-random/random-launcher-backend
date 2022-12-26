package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"

	"github.com/ch-random/random-launcher-backend/domain"
)

type userUsecase struct {
	ur     domain.UserRepository
	timeout time.Duration
}

func NewUserUsecase(
	ur domain.UserRepository,
	timeout time.Duration,
) domain.UserUsecase {
	return &userUsecase{
		ur,
		timeout,
	}
}

func (uu *userUsecase) Fetch(c context.Context, cursor string, numString string) (us []domain.User, nextCursor string, err error) {
	_, cancel := context.WithTimeout(c, uu.timeout)
	defer cancel()

	us, nextCursor, err = uu.ur.Fetch(cursor, numString)
	if err != nil {
		return nil, "", err
	}
	return
}

func fillNewUserDetails(u *domain.User) (*domain.User, error) {
	uid := uuid.New()
	log.Printf("uid: %v", uid)
	u.ID = uid
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	return u, nil
}
func (uu *userUsecase) Insert(c context.Context, u *domain.User) (err error) {
	ctx, cancel := context.WithTimeout(c, uu.timeout)
	defer cancel()

	// Check for id conflicts
	_, err = uu.GetByID(ctx, u.ID)
	log.Printf("uu.go %v", err)
	if err == nil {
		return domain.ErrConflict
	} else if err != domain.ErrNotFound {
		return err
	}

	u, err = fillNewUserDetails(u)
	if err != nil {
		return err
	}
	log.Printf("u: %v", u)
	return uu.ur.Insert(u)
}

func (uu *userUsecase) GetByID(c context.Context, id uuid.UUID) (u domain.User, err error) {
	_, cancel := context.WithTimeout(c, uu.timeout)
	defer cancel()

	u, err = uu.ur.GetByID(id)
	if err != nil {
		return domain.User{}, err
	}
	return
}

func (uu *userUsecase) Update(c context.Context, ac *domain.User) error {
	_, cancel := context.WithTimeout(c, uu.timeout)
	defer cancel()

	ac.UpdatedAt = time.Now()
	return uu.ur.Update(ac)
}
