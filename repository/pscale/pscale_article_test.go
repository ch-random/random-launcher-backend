package pscale_test

// import (
// 	"testing"
// 	"time"

// 	"github.com/stretchr/testify/assert"
// 	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"

// 	"github.com/ch-random/random-launcher-backend/repository"
// 	"github.com/ch-random/random-launcher-backend/repository/pscale"
// 	"github.com/ch-random/random-launcher-backend/domain"
// )

// func TestFetch(t *testing.T) {
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 	}

// 	mockArticles := []domain.Article{
// 		domain.Article{
// 			ID: 1, Title: "title 1", Body: "content 1",
// 			User: domain.User{ID: 1}, UpdatedAt: time.Now(), CreatedAt: time.Now(),
// 		},
// 		domain.Article{
// 			ID: 2, Title: "title 2", Body: "content 2",
// 			User: domain.User{ID: 1}, UpdatedAt: time.Now(), CreatedAt: time.Now(),
// 		},
// 	}

// 	rows := sqlmock.NewRows([]string{"id", "title", "content", "user_id", "updated_at", "created_at"}).
// 		AddRow(mockArticles[0].ID, mockArticles[0].Title, mockArticles[0].Body,
// 			mockArticles[0].User.ID, mockArticles[0].UpdatedAt, mockArticles[0].CreatedAt).
// 		AddRow(mockArticles[1].ID, mockArticles[1].Title, mockArticles[1].Body,
// 			mockArticles[1].User.ID, mockArticles[1].UpdatedAt, mockArticles[1].CreatedAt)

// 	query := "SELECT id,title,content, user_id, updated_at, created_at FROM article WHERE created_at > \\? ORDER BY created_at LIMIT \\?"

// 	mock.ExpectQuery(query).WillReturnRows(rows)
// 	a := pscale.NewPscaleArticleRepository(db)
// 	cursor := repository.EncodeCursor(mockArticles[1].CreatedAt)
// 	num := int(2)
// 	list, nextCursor, err := a.Fetch(cursor, num)
// 	assert.NotEmpty(t, nextCursor)
// 	assert.NoError(t, err)
// 	assert.Len(t, list, 2)
// }

// func TestGetByID(t *testing.T) {
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 	}

// 	rows := sqlmock.NewRows([]string{"id", "title", "content", "user_id", "updated_at", "created_at"}).
// 		AddRow(1, "title 1", "Body 1", 1, time.Now(), time.Now())

// 	query := "SELECT id,title,content, user_id, updated_at, created_at FROM article WHERE ID = \\?"

// 	mock.ExpectQuery(query).WillReturnRows(rows)
// 	a := pscale.NewPscaleArticleRepository(db)

// 	num := uint(5)
// 	anArticle, err := a.GetByID(num)
// 	assert.NoError(t, err)
// 	assert.NotNil(t, anArticle)
// }

// func TestInsert(t *testing.T) {
// 	now := time.Now()
// 	ar := &domain.Article{
// 		Title:     "Judul",
// 		Body:   "Body",
// 		CreatedAt: now,
// 		UpdatedAt: now,
// 		User: domain.User{
// 			ID:   1,
// 			Name: "Iman Tumorang",
// 		},
// 	}
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 	}

// 	query := "INSERT  article SET title=\\? , content=\\? , user_id=\\?, updated_at=\\? , created_at=\\?"
// 	prep := mock.ExpectPrepare(query)
// 	prep.ExpectExec().WithArgs(ar.Title, ar.Body, ar.User.ID, ar.CreatedAt, ar.UpdatedAt).WillReturnResult(sqlmock.NewResult(12, 1))

// 	a := pscale.NewPscaleArticleRepository(db)

// 	err = a.Insert(ar)
// 	assert.NoError(t, err)
// 	assert.Equal(t, uint(12), ar.ID)
// }

// func TestGetByTitle(t *testing.T) {
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 	}

// 	rows := sqlmock.NewRows([]string{"id", "title", "content", "user_id", "updated_at", "created_at"}).
// 		AddRow(1, "title 1", "Body 1", 1, time.Now(), time.Now())

// 	query := "SELECT id,title,content, user_id, updated_at, created_at FROM article WHERE title = \\?"

// 	mock.ExpectQuery(query).WillReturnRows(rows)
// 	a := pscale.NewPscaleArticleRepository(db)

// 	title := "title 1"
// 	anArticle, err := a.GetByTitle(title)
// 	assert.NoError(t, err)
// 	assert.NotNil(t, anArticle)
// }

// func TestDelete(t *testing.T) {
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 	}

// 	query := "DELETE FROM article WHERE id = \\?"

// 	prep := mock.ExpectPrepare(query)
// 	prep.ExpectExec().WithArgs(12).WillReturnResult(sqlmock.NewResult(12, 1))

// 	a := pscale.NewPscaleArticleRepository(db)

// 	num := uint(12)
// 	err = a.Delete(num)
// 	assert.NoError(t, err)
// }

// func TestUpdate(t *testing.T) {
// 	now := time.Now()
// 	ar := &domain.Article{
// 		ID:        12,
// 		Title:     "Judul",
// 		Body:   "Body",
// 		CreatedAt: now,
// 		UpdatedAt: now,
// 		User: domain.User{
// 			ID:   1,
// 			Name: "Iman Tumorang",
// 		},
// 	}

// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 	}

// 	query := "UPDATE article set title=\\?, content=\\?, user_id=\\?, updated_at=\\? WHERE ID = \\?"

// 	prep := mock.ExpectPrepare(query)
// 	prep.ExpectExec().WithArgs(ar.Title, ar.Body, ar.User.ID, ar.UpdatedAt, ar.ID).WillReturnResult(sqlmock.NewResult(12, 1))

// 	a := pscale.NewPscaleArticleRepository(db)

// 	err = a.Update(ar)
// 	assert.NoError(t, err)
// }
