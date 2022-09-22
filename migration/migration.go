package migration

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"

	"github.com/ch-random/random-launcher-backend/domain"
	"github.com/ch-random/random-launcher-backend/migration/version"
	"github.com/ch-random/random-launcher-backend/utils"
)

const tableName = "migrations"

var (
	Versions = []func(*gorm.DB) error{
		// version.VyyyyMMdd, // change log
		version.V20220922, // 新規作成
	}

	// All tables in the latest schema
	AllTables = []interface{}{
		&domain.User{},
		&domain.Article{},
		&domain.ArticleGameContent{},
		&domain.ArticleOwner{},
		&domain.ArticleTag{},
		&domain.ArticleComment{},
		&domain.ArticleImageURL{},
	}
)

// https://github.com/go-gormigrate/gormigrate
func Migrate(db *gorm.DB) (inited bool, err error) {
	var ms = []*gormigrate.Migration{}
	for _, v := range Versions {
		ms = append(ms, &gormigrate.Migration{
			ID:      utils.GetFuncName(v),
			Migrate: v,
		})
	}
	gm := gormigrate.New(db, &gormigrate.Options{
		TableName:                 tableName,
		IDColumnName:              "id",
		IDColumnSize:              128,
		UseTransaction:            false,
		ValidateUnknownMigrations: true,
	}, ms)

	gm.InitSchema(func(db *gorm.DB) error {
		inited = true
		return db.AutoMigrate(AllTables...)
	})
	// `Migrate` executes all migrations that did not run yet.
	err = gm.Migrate()
	return
}

func DropAllTables(db *gorm.DB) error {
	m := db.Migrator()
	if err := m.DropTable(AllTables...); err != nil {
		return err
	}
	return m.DropTable(tableName)
}
