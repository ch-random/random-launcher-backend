// https://github.com/planetscale/golang-example/blob/main/main.go
// REST API設計者のための有名APIのURL例
// https://zenn.dev/yu1ro/articles/4c73274383b676
package httpserver

import (
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	validator "gopkg.in/go-playground/validator.v9"
	"gorm.io/gorm"

	"github.com/ch-random/random-launcher-backend/domain"
)

type ResponseError struct {
	Message string `json:"message"`
}

type ArticleHandler struct {
	ArticleUsecase domain.ArticleUsecase
}

func NewArticleHandler(db *gorm.DB, articleUsecase domain.ArticleUsecase) *echo.Echo {
	h := &ArticleHandler{ArticleUsecase: articleUsecase}
	e := echo.New()

	e.GET("/articles", h.FetchArticles)
	e.POST("/articles", h.InsertArticle)
	e.PUT("/articles", h.UpdateArticle)

	e.GET("/articles/:id", h.GetByID)
	e.DELETE("/articles/:id", h.Delete)

	return e
}

func articleValid(ar *domain.Article) (bool, error) {
	validate := validator.New()
	if err := validate.Struct(ar); err != nil {
		return false, err
	}
	return true, nil
}

func (h *ArticleHandler) FetchArticles(c echo.Context) error {
	numString := c.QueryParam("num")
	cursor := c.QueryParam("cursor")
	ctx := c.Request().Context()

	articles, nextCursor, err := h.ArticleUsecase.Fetch(ctx, cursor, numString)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	c.Response().Header().Set(`X-Cursor`, nextCursor)
	return c.JSON(http.StatusOK, articles)
}
func (h *ArticleHandler) InsertArticle(c echo.Context) (err error) {
	var article domain.Article
	if err = c.Bind(&article); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := articleValid(&article); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()
	if err = h.ArticleUsecase.Insert(ctx, &article); err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusCreated, article)
}
func (h *ArticleHandler) UpdateArticle(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.ErrBadParamInput.Error())
	}

	var ar domain.Article
	if err := c.Bind(&ar); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, domain.ErrBadRequestBodyInput.Error())
	}
	ar.ID = uint(id)

	ctx := c.Request().Context()
	err = h.ArticleUsecase.Update(ctx, &ar)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, ar)
}

func (h *ArticleHandler) GetByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.ErrBadParamInput.Error())
	}

	ctx := c.Request().Context()
	article, err := h.ArticleUsecase.GetByID(ctx, uint(id))
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, article)
}
func (h *ArticleHandler) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, domain.ErrBadParamInput.Error())
	}

	ctx := c.Request().Context()
	if err = h.ArticleUsecase.Delete(ctx, uint(id)); err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	log.Println(err)
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
