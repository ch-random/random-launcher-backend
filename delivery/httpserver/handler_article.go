// https://github.com/planetscale/golang-example/blob/main/main.go
// REST API設計者のための有名APIのURL例
// https://zenn.dev/yu1ro/articles/4c73274383b676
package httpserver

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo"
	validator "gopkg.in/go-playground/validator.v9"

	"github.com/ch-random/random-launcher-backend/domain"
)

func articleValid(ar *domain.Article) (bool, error) {
	validate := validator.New()
	if err := validate.Struct(ar); err != nil {
		return false, err
	}
	return true, nil
}

func (h *HTTPHandler) FetchArticles(c echo.Context) error {
	numString := c.QueryParam("num")
	cursor := c.QueryParam("cursor")
	ctx := c.Request().Context()

	ars, nextCursor, err := h.ArticleUsecase.Fetch(ctx, cursor, numString)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	c.Response().Header().Set(`X-Cursor`, nextCursor)
	return c.JSON(http.StatusOK, ars)
}
func (h *HTTPHandler) InsertArticle(c echo.Context) error {
	var ar domain.Article
	if err := c.Bind(&ar); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := articleValid(&ar); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()
	if err := h.ArticleUsecase.Insert(ctx, &ar); err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusCreated, ar)
}

func (h *HTTPHandler) GetArticleByID(c echo.Context) error {
	idString := c.Param("id")
	id, err := uuid.Parse(idString)
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.ErrBadParamInput.Error())
	}

	ctx := c.Request().Context()
	article, err := h.ArticleUsecase.GetByID(ctx, id)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, article)
}
func (h *HTTPHandler) UpdateArticle(c echo.Context) error {
	idString := c.Param("id")
	id, err := uuid.Parse(idString)
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.ErrBadParamInput.Error())
	}

	var ar domain.Article
	if err := c.Bind(&ar); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, domain.ErrBadRequestBodyInput.Error())
	}
	ar.ID = id

	ctx := c.Request().Context()
	err = h.ArticleUsecase.Update(ctx, &ar)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, ar)
}
func (h *HTTPHandler) DeleteArticle(c echo.Context) error {
	idString := c.Param("id")
	id, err := uuid.Parse(idString)
	if err != nil {
		return c.JSON(http.StatusNotFound, domain.ErrBadParamInput.Error())
	}

	ctx := c.Request().Context()
	if err = h.ArticleUsecase.Delete(ctx, id); err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}
