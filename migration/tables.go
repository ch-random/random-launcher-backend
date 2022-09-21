package migration

import "github.com/ch-random/random-launcher-backend/domain"

// All tables in the latest schema
func AllTables() []interface{} {
	return []interface{}{
		&domain.User{},
		&domain.Article{},
		&domain.ArticleGameContent{},
		&domain.ArticleOwner{},
		&domain.ArticleTag{},
		&domain.ArticleComment{},
		&domain.ArticleImageURL{},
	}
}
