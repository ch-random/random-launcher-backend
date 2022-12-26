package httpserver

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo"
	"github.com/rs/zerolog/log"

	"github.com/ch-random/random-launcher-backend/domain"
)

func (h *httpHandler) FetchUsers(c echo.Context) error {
	cursor := c.QueryParam("cursor")
	numString := c.QueryParam("num")
	ctx := c.Request().Context()

	us, nextCursor, err := h.UserUsecase.Fetch(ctx, cursor, numString)
	if err != nil {
		return c.JSON(getStatusCode(err), getResponseError(err))
	}
	c.Response().Header().Set(`X-Cursor`, nextCursor)
	return c.JSON(http.StatusOK, us)
}
func (h *httpHandler) InsertUser(c echo.Context) error {
	var u domain.User
	if err := c.Bind(&u); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, getResponseError(err))
	}

	if ok, err := valid(&u); !ok {
		return c.JSON(http.StatusBadRequest, getResponseError(err))
	}

	ctx := c.Request().Context()
	if err := h.UserUsecase.Insert(ctx, &u); err != nil {
		log.Printf("hu.go %v", err)
		return c.JSON(getStatusCode(err), getResponseError(err))
	}
	return c.JSON(http.StatusCreated, u)
}

func (h *httpHandler) UpdateUser(c echo.Context) error {
	var u domain.User
	if err := c.Bind(&u); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, getResponseError(domain.ErrBadRequestBodyInput))
	}

	// verify
	idString := c.Param("id")
	id, err := uuid.Parse(idString)
	if err != nil {
		return c.JSON(http.StatusBadRequest, getResponseError(domain.ErrBadParamInput))
	}
	if u.ID != id {
		return c.JSON(http.StatusBadRequest, getResponseError(domain.ErrBadParamInput))
	}


	ctx := c.Request().Context()
	err = h.UserUsecase.Update(ctx, &u)
	if err != nil {
		return c.JSON(getStatusCode(err), getResponseError(err))
	}
	return c.JSON(http.StatusOK, u)
}
