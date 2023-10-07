package httpserver

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/rs/zerolog/log"

	"github.com/ch-random/random-launcher-backend/config"
	"github.com/ch-random/random-launcher-backend/domain"
	"github.com/ch-random/random-launcher-backend/middleware/httpheader"
	"github.com/ch-random/random-launcher-backend/migration"
	"github.com/ch-random/random-launcher-backend/repository"
	"github.com/ch-random/random-launcher-backend/repository/pscale"
	"github.com/ch-random/random-launcher-backend/usecase"
)

type ResponseError struct {
	Message string `json:"message"`
	Version string `json:"version"`
}
type httpHandler struct {
	UserUsecase           domain.UserUsecase
	ArticleUsecase        domain.ArticleUsecase
	ArticleCommentUsecase domain.ArticleCommentUsecase
}

func NewHandler() *echo.Echo {
	e := echo.New()

	// CORS, SUPA_ANON_KEY
	httpHeaderHandler := httpheader.NewHTTPHeaderHandler()
	e.Use(httpHeaderHandler.HandleHTTPHeader)

	db, err := pscale.GetDB()
	if err != nil {
		panic("failed to connect to PlanetScale")
	}

	ur := pscale.NewUserRepository(db)
	ar := pscale.NewArticleRepository(db)
	agc := pscale.NewArticleGameContentRepository(db)
	aor := pscale.NewArticleOwnerRepository(db)
	atr := pscale.NewArticleTagRepository(db)
	acr := pscale.NewArticleCommentRepository(db)
	aiur := pscale.NewArticleImageURLRepository(db)
	timeout := config.Timeout
	h := &httpHandler{
		UserUsecase: usecase.NewUserUsecase(ur, timeout),
		ArticleUsecase: usecase.NewArticleUsecase(ur, ar, agc, aor, atr, acr, aiur, timeout),
		ArticleCommentUsecase: usecase.NewArticleCommentUsecase(acr, timeout),
	}

	// /
	e.GET("/", h.Index)

	// /users
	e.GET("/users", h.FetchUsers)
	e.POST("/users", h.InsertUser)
	// /users/:id
	e.PUT("/users/:id", h.UpdateUser)

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
	e.DELETE("/comments/:id", h.DeleteArticleCommentByID)
	// /comments/article/:id
	e.GET("/comments/article/:id", h.GetArticleCommentsByArticleID)
	e.DELETE("/comments/article/:id", h.DeleteCommentByArticleID)
	return e
}

func valid(u interface{}) (bool, error) {
	v := repository.NewValidator()
	if err := v.Struct(u); err != nil {
		return false, err
	}
	return true, nil
}

func (h *httpHandler) Index(c echo.Context) (err error) {
	err = domain.ErrNotFound
	return c.JSON(getStatusCode(err), getResponseError(err))
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
