package httpserver

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/rs/zerolog/log"

	"github.com/ch-random/random-launcher-backend/configs"
	"github.com/ch-random/random-launcher-backend/domain"
	"github.com/ch-random/random-launcher-backend/middleware/cors"
	"github.com/ch-random/random-launcher-backend/migration"
	"github.com/ch-random/random-launcher-backend/repository/pscale"
	"github.com/ch-random/random-launcher-backend/usecase"
)

type ResponseError struct {
	Message string `json:"message"`
	Version string `json:"version"`
}

type HTTPHandler struct {
	ArticleUsecase        domain.ArticleUsecase
	ArticleCommentUsecase domain.ArticleCommentUsecase
}

func NewHandler() *echo.Echo {
	e := echo.New()

	corsHandler := cors.NewCORSHandler()
	e.Use(corsHandler.HandleCORS)

	h := &HTTPHandler{ArticleUsecase: newArticleUsecase()}

	// /
	e.GET("/", h.Index)

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

func (h *HTTPHandler) Index(c echo.Context) (err error) {
	err = domain.ErrNotFound
	return c.JSON(getStatusCode(err), getResponseError(err))
}

func newArticleUsecase() domain.ArticleUsecase {
	db, err := pscale.GetDB()
	if err != nil {
		log.Warn().Err(err).Msg("failed to connect to PlanetScale")
	}

	ur := pscale.NewUserRepository(db)
	ar := pscale.NewArticleRepository(db)
	agc := pscale.NewArticleGameContentRepository(db)
	aor := pscale.NewArticleOwnerRepository(db)
	atr := pscale.NewArticleTagRepository(db)
	acr := pscale.NewArticleCommentRepository(db)
	aiur := pscale.NewArticleImageURLRepository(db)
	timeout := configs.Timeout
	au := usecase.NewArticleUsecase(
		ur,
		ar,
		agc,
		aor,
		atr,
		acr,
		aiur,
		timeout,
	)
	return au
}

func getResponseError(err error) ResponseError {
	vs := migration.GetVersions()
	return ResponseError{
		Message: err.Error(),
		Version: vs[len(vs)-1],
	}
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	log.Info().Err(err).Msg("getStatusCode")
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
