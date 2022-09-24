// https://github.com/planetscale/golang-example/blob/main/main.go
// REST API設計者のための有名APIのURL例
// https://zenn.dev/yu1ro/articles/4c73274383b676
package httpserver

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/ch-random/random-launcher-backend/domain"
	"github.com/ch-random/random-launcher-backend/repository"
)

func articleCommentValid(ac *domain.ArticleComment) (bool, error) {
	v := repository.NewValidator()
	if err := v.Struct(ac); err != nil {
		return false, err
	}
	return true, nil
}

func (h *HTTPHandler) InsertArticleComment(c echo.Context) error {
	var ac domain.ArticleComment
	if err := c.Bind(&ac); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := articleCommentValid(&ac); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()
	if err := h.ArticleCommentUsecase.Insert(ctx, &ac); err != nil {
		return c.JSON(getStatusCode(err), getResponseError(err))
	}
	return c.JSON(http.StatusCreated, ac)
}

func (h *HTTPHandler) UpdateArticleComment(c echo.Context) error {
	return nil
}
func (h *HTTPHandler) DeleteeArticleCommentByID(c echo.Context) error {
	return nil
}

func (h *HTTPHandler) GetArticleCommentsByArticleID(c echo.Context) error {
	return nil
}
func (h *HTTPHandler) DeleteCommentByArticleID(c echo.Context) error {
	return nil
}
