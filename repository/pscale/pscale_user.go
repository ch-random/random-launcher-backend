package pscale

import (
	"strconv"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"

	"github.com/ch-random/random-launcher-backend/domain"
	"github.com/ch-random/random-launcher-backend/repository"
)

type pscaleUserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	// return &pscaleUserRepository{db.Model(&domain.User{})}
	return &pscaleUserRepository{db}
}

func (userRepo *pscaleUserRepository) Fetch(cursor string, numString string) (us []domain.User, nextCursor string, err error) {
	db := userRepo.db

	if cursor != "" {
		decodedCursor, err := repository.DecodeCursor(cursor)
		log.Printf("decodedCursor: %v", decodedCursor)
		if err != nil && cursor != "" {
			return nil, "", domain.ErrBadParamInput
		}
		db = db.Where("created_at > ?", decodedCursor).Order("created_at")
		if err = db.Error; err != nil {
			return nil, "", err
		}
	}

	num, err := strconv.Atoi(numString)
	if err == nil {
		log.Printf("num: %v", num)
		db = db.Limit(num)
	}

	// Preloading (Eager Loading) https://gorm.io/ja_JP/docs/preload.html
	// Join Preload は、one-to-one (has one, belongs to) でのみ動作します。
	// Preloading https://qiita.com/tsubasaozawa/items/ac5a8a515fe4f7a139b0
	// Nested Preloading: db.Preload("ArticleOwner.User")
	// db.Preload(clause.Associations) により全データを preload できる。
	// 不要データを含むなら、db.Preload(""), `gorm:"PRELOAD:false"` として高速化。
	db = db.Preload(clause.Associations).Begin().Find(&us)
	if err = db.Error; err != nil {
		return
	}
	log.Printf("us: %v", us)

	// if len(us) > 0 && len(us) == num {
	// 	nextCursor = repository.EncodeCursor(us[len(us)-1].CreatedAt)
	// }
	return
}

func (userRepo *pscaleUserRepository) GetByID(id uuid.UUID) (user domain.User, err error) {
	// query := "SELECT id, name, created_at, updated_at FROM user WHERE id=?"
	// return userRepo.getOne(ctx, query, id)
	if db := userRepo.db.Where("id = ?", id).First(&user); db.Error != nil {
		err = domain.ErrNotFound
	}
	return
}

func (userRepo *pscaleUserRepository) Insert(u *domain.User) (err error) {
	db := userRepo.db.Create(u)
	if err = db.Error; err != nil {
		db.Rollback()
		return
	}
	return
}

func (userRepo *pscaleUserRepository) Update(u *domain.User) (err error) {
	db := userRepo.db.Where("id = ?", u.ID).Updates(u)
	if err = db.Error; err != nil {
		return
	}

	if affected := db.RowsAffected; affected != 1 {
		err = domain.ErrRowsAffectedNotOne
	}
	return
}
