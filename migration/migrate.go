package migration

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"

	"github.com/ch-random/random-launcher-backend/migration/version"
)

var ms = []*gormigrate.Migration{
	// version.VyyyyMMdd(), // change log
	version.V20220921(), // 新規作成
}

func Migrate(db *gorm.DB) (init bool, err error) {
	gm := gormigrate.New(db, &gormigrate.Options{
		TableName:                 "migrations",
		IDColumnName:              "id",
		IDColumnSize:              128,
		UseTransaction:            false,
		ValidateUnknownMigrations: true,
	}, ms)

	gm.InitSchema(func(db *gorm.DB) error {
		init = true
		return db.AutoMigrate(AllTables()...)
	})
	err = gm.Migrate()
	return
}

func DropAll(db *gorm.DB) error {
	if err := db.Migrator().DropTable(AllTables()...); err != nil {
		return err
	}
	return db.Migrator().DropTable("migrations")
}
