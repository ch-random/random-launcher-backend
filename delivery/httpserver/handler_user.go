package httpserver

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo"

	"github.com/ch-random/random-launcher-backend/domain"
	// "github.com/ch-random/random-launcher-backend/repository"
)

// func userValid(u *domain.User) (bool, error) {
// 	v := repository.NewValidator()
// 	if err := v.Struct(u); err != nil {
// 		return false, err
// 	}
// 	return true, nil
// }

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
