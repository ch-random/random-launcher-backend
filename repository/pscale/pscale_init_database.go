package pscale

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/rs/zerolog/log"
)

func GetDB() (db *gorm.DB, err error) {
	dsn := os.Getenv("DSN")
	if dsn == "" {
		log.Fatal().Msg("DSN environment variable is not set")
	}

	// Connect to PlanetScale database using DSN environment variable.
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		return nil, err
	}
	return
}
