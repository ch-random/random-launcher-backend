// https://github.com/planetscale/golang-example/blob/main/main.go
// REST API設計者のための有名APIのURL例
// https://zenn.dev/yu1ro/articles/4c73274383b676
package httpserver

import (
	// "log"
	"net/http"
	// "strconv"

	"github.com/labstack/echo"
	validator "gopkg.in/go-playground/validator.v9"
	// "gorm.io/gorm"

	"github.com/ch-random/random-launcher-backend/domain"
)

func articleCommentValid(ac *domain.ArticleComment) (bool, error) {
	validate := validator.New()
	if err := validate.Struct(ac); err != nil {
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
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
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
