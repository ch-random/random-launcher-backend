package httpserver

import (
	"net/http"

	"github.com/google/uuid"
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

func (h *httpHandler) InsertArticleComment(c echo.Context) error {
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

func (h *httpHandler) UpdateArticleComment(c echo.Context) error {
	var ac domain.ArticleComment
	if err := c.Bind(&ac); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, getResponseError(domain.ErrBadRequestBodyInput))
	}
	// ac.ID = id

	ctx := c.Request().Context()
	err := h.ArticleCommentUsecase.Update(ctx, &ac)
	if err != nil {
		return c.JSON(getStatusCode(err), getResponseError(err))
	}
	return c.JSON(http.StatusOK, ac)
}
func (h *httpHandler) DeleteArticleCommentByID(c echo.Context) error {
	idString := c.Param("id")
	id, err := uuid.Parse(idString)
	if err != nil {
		return c.JSON(http.StatusNotFound, getResponseError(domain.ErrBadParamInput))
	}

	ctx := c.Request().Context()
	if err = h.ArticleCommentUsecase.Delete(ctx, id); err != nil {
		return c.JSON(getStatusCode(err), getResponseError(err))
	}
	return c.NoContent(http.StatusNoContent)
}

func (h *httpHandler) GetArticleCommentsByArticleID(c echo.Context) error {
	idString := c.Param("id")
	id, err := uuid.Parse(idString)
	if err != nil {
		return c.JSON(http.StatusBadRequest, getResponseError(domain.ErrBadParamInput))
	}

	ctx := c.Request().Context()
	acs, err := h.ArticleCommentUsecase.GetByArticleID(ctx, id)
	if err != nil {
		return c.JSON(getStatusCode(err), getResponseError(err))
	}
	return c.JSON(http.StatusOK, acs)
}
func (h *httpHandler) DeleteCommentByArticleID(c echo.Context) error {
	idString := c.Param("id")
	id, err := uuid.Parse(idString)
	if err != nil {
		return c.JSON(http.StatusNotFound, getResponseError(domain.ErrBadParamInput))
	}

	ctx := c.Request().Context()
	if err = h.ArticleCommentUsecase.Delete(ctx, id); err != nil {
		return c.JSON(getStatusCode(err), getResponseError(err))
	}
	return c.NoContent(http.StatusNoContent)
}
