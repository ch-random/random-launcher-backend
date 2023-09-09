package domain

import (
	"gorm.io/gorm"
)

type Migration struct {
	Migrate  func(db *gorm.DB) error
	Rollback func(db *gorm.DB) error
}
