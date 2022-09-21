package pscale

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetDB() (db *gorm.DB, err error) {
	// Connect to PlanetScale database using DSN environment variable.
	db, err = gorm.Open(mysql.Open(os.Getenv("DSN")), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		return nil, err
	}
	return
}
