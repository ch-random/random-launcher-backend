package pscale

import (
	"gorm.io/gorm"

	"github.com/ch-random/random-launcher-backend/domain"
)

type pscaleUserRepo struct {
	db *gorm.DB
}

// NewPscaleUserRepository will create an implementation of user.Repository
func NewPscaleUserRepository(db *gorm.DB) domain.UserRepository {
	return &pscaleUserRepo{db}
}

func (userRepo *pscaleUserRepo) GetByID(id uint) (user domain.User, err error) {
	// query := "SELECT id, name, created_at, updated_at FROM user WHERE id=?"
	// return userRepo.getOne(ctx, query, id)
	if result := userRepo.db.Where("id = ?", id).First(&user); result.Error != nil {
		err = domain.ErrNotFound
	}
	return
}
