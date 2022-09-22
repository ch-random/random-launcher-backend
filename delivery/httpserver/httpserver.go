package httpserver

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"

	"github.com/ch-random/random-launcher-backend/domain"
)

type ResponseError struct {
	Message string `json:"message"`
}

type HTTPHandler struct {
	ArticleUsecase domain.ArticleUsecase
	ArticleCommentUsecase domain.ArticleCommentUsecase
}

func NewHandler(db *gorm.DB, articleUsecase domain.ArticleUsecase) *echo.Echo {
	h := &HTTPHandler{ArticleUsecase: articleUsecase}
	e := echo.New()

	// /articles
	e.GET("/articles", h.FetchArticles)
	e.POST("/articles", h.InsertArticle)
	// /articles/:id
	e.GET("/articles/:id", h.GetArticleByID)
	e.PUT("/articles/:id", h.UpdateArticle)
	e.DELETE("/articles/:id", h.DeleteArticle)

	// /comments
	e.POST("/comments", h.InsertArticleComment)
	// /comments/:id
	e.PUT("/comments/:id", h.UpdateArticleComment)
	e.DELETE("/comments/:id", h.DeleteeArticleCommentByID)
	// /comments/article/:id
	e.GET("/comments/article/:id", h.GetArticleCommentsByArticleID)
	e.DELETE("/comments/article/:id", h.DeleteCommentByArticleID)
	return e
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	log.Print(err)
	switch err {
	case domain.ErrInternalServerError:
		return http.StatusInternalServerError
	case domain.ErrNotFound:
		return http.StatusNotFound
	case domain.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
