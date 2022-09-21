package version

import (
	"gorm.io/gorm"
	gormigrate "github.com/go-gormigrate/gormigrate/v2"
)

// https://github.com/go-gormigrate/gormigrate
func V20220921() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "20220921",
		Migrate: func(db *gorm.DB) error {
			if err := db.AutoMigrate(&v20220921Something{}); err != nil {
				return err
			}
			return db.Exec("").Error
		},
	}
}

type v20220921Something struct {}
