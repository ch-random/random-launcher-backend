package pscale

import (
	"os"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDB() (db *gorm.DB, err error) {
	dsn := os.Getenv("DSN")
	if dsn == "" {
		log.Fatal().Msg("DSN environment variable is not set")
	}

	// https://gorm.io/ja_JP/docs/migration.html
	// Connect to PlanetScale database using DSN environment variable
	// db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		return nil, err
	}
	return
}
