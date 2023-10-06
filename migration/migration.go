// https://github.com/go-gormigrate/gormigrate
// gopkg.in/gormigrate.v2
package migration

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"

	"github.com/ch-random/random-launcher-backend/domain"
	"github.com/ch-random/random-launcher-backend/migration/version"
	"github.com/ch-random/random-launcher-backend/utils"
)

const migrationTableName = "migrations"

var (
	migrations = []func() domain.Migration{
		// version.VyyyyMMdd(), // change log
		version.V20220922, // 新規作成
		// version.V20230830, // Article.EventId を追加
	}

	GormigrateOptions = &gormigrate.Options{
		TableName:                 migrationTableName,
		IDColumnName:              "id",
		IDColumnSize:              128,
		UseTransaction:            false,
		ValidateUnknownMigrations: true,
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

func GetVersions() []string {
	var vs []string
	for _, m := range migrations {
		vs = append(vs, utils.GetFuncName(m))
	}
	return vs
}

func Migrate(db *gorm.DB) (inited bool, err error) {
	var ms = []*gormigrate.Migration{}
	for _, m := range migrations {
		ms = append(ms, &gormigrate.Migration{
			ID:       utils.GetFuncName(m),
			Migrate:  m().Migrate,
			Rollback: m().Rollback,
		})
	}
	gm := gormigrate.New(db, GormigrateOptions, ms)

	gm.InitSchema(func(db *gorm.DB) error {
		inited = true
		return db.AutoMigrate(AllTables...)
	})

	// `Migrate` executes all migrations that did not run yet
	err = gm.Migrate()
	return
}

func DropAllTables(db *gorm.DB) error {
	m := db.Migrator()
	if err := m.DropTable(AllTables...); err != nil {
		return err
	}
	return m.DropTable(migrationTableName)
}
