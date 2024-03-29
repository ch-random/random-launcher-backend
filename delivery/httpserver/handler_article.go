// https://github.com/planetscale/golang-example/blob/main/main.go
// REST API設計者のための有名APIのURL例
// https://zenn.dev/yu1ro/articles/4c73274383b676
package httpserver

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo"
	// "github.com/rs/zerolog/log"

	"github.com/ch-random/random-launcher-backend/domain"
)

func (h *httpHandler) FetchArticles(c echo.Context) error {
	cursor := c.QueryParam("cursor")
	numString := c.QueryParam("num")
	ctx := c.Request().Context()

	ars, nextCursor, err := h.ArticleUsecase.Fetch(ctx, cursor, numString)
	if err != nil {
		return c.JSON(getStatusCode(err), getResponseError(err))
	}
	c.Response().Header().Set(`X-Cursor`, nextCursor)
	return c.JSON(http.StatusOK, ars)
}
func (h *httpHandler) InsertArticle(c echo.Context) error {
	var ar domain.Article
	if err := c.Bind(&ar); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, getResponseError(err))
	}

	if ok, err := valid(&ar); !ok {
		return c.JSON(http.StatusBadRequest, getResponseError(err))
	}

	ctx := c.Request().Context()
	if err := h.ArticleUsecase.Insert(ctx, &ar); err != nil {
		return c.JSON(getStatusCode(err), getResponseError(err))
	}
	return c.JSON(http.StatusCreated, ar)
}

func (h *httpHandler) GetArticleByID(c echo.Context) error {
	idString := c.Param("id")
	id, err := uuid.Parse(idString)
	if err != nil {
		return c.JSON(http.StatusBadRequest, getResponseError(domain.ErrBadParamInput))
	}

	ctx := c.Request().Context()
	ar, err := h.ArticleUsecase.GetByID(ctx, id)
	if err != nil {
		return c.JSON(getStatusCode(err), getResponseError(err))
	}
	return c.JSON(http.StatusOK, ar)
}
func (h *httpHandler) UpdateArticle(c echo.Context) error {
	var ar domain.Article
	if err := c.Bind(&ar); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, getResponseError(domain.ErrBadRequestBodyInput))
	}

	// verify
	// https://zenn.dev/skanehira/articles/2020-09-19-go-echo-bind-tips
	idString := c.Param("id")
	id, err := uuid.Parse(idString)
	if err != nil {
		return c.JSON(http.StatusBadRequest, getResponseError(domain.ErrBadParamInput))
	}
	if ar.ID != id {
		return c.JSON(http.StatusBadRequest, getResponseError(domain.ErrBadParamInput))
	}

	ctx := c.Request().Context()
	err = h.ArticleUsecase.Update(ctx, &ar)
	if err != nil {
		return c.JSON(getStatusCode(err), getResponseError(err))
	}
	return c.JSON(http.StatusOK, ar)
}
func (h *httpHandler) DeleteArticle(c echo.Context) error {
	idString := c.Param("id")
	id, err := uuid.Parse(idString)
	if err != nil {
		return c.JSON(http.StatusNotFound, getResponseError(domain.ErrBadParamInput))
	}

	ctx := c.Request().Context()
	if err = h.ArticleUsecase.Delete(ctx, id); err != nil {
		return c.JSON(getStatusCode(err), getResponseError(err))
	}
	return c.NoContent(http.StatusNoContent)
}
