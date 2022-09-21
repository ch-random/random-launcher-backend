package pscale

import (
	"gorm.io/gorm"
	"github.com/google/uuid"

	"github.com/ch-random/random-launcher-backend/domain"
)

type pscaleUserRepo struct {
	db *gorm.DB
}

// NewUserRepository will create an implementation of user.Repository
func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &pscaleUserRepo{db}
}

func (userRepo *pscaleUserRepo) GetByID(id uuid.UUID) (user domain.User, err error) {
	// query := "SELECT id, name, created_at, updated_at FROM user WHERE id=?"
	// return userRepo.getOne(ctx, query, id)
	if db := userRepo.db.Where("id = ?", id).First(&user); db.Error != nil {
		err = domain.ErrNotFound
	}
	return
}
