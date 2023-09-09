package version

import (
	"gorm.io/gorm"

	"github.com/ch-random/random-launcher-backend/domain"
)

func V20220922() domain.Migration {
	migrate := func(db *gorm.DB) (err error) {
		return
	}
	rollback := func(db *gorm.DB) (err error) {
		return
	}
	return domain.Migration{
		Migrate:  migrate,
		Rollback: rollback,
	}
}
