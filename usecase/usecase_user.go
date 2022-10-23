package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"

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
