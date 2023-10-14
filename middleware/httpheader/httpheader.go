package httpheader

import (
	"os"

	"github.com/labstack/echo"
)

type HTTPHeaderHandler struct{}

func (m *HTTPHeaderHandler) HandleHTTPHeader(handle echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// CORS
		c.Response().Header().Set(echo.HeaderAccessControlAllowOrigin, "*")
		// Content-Type
		// c.Response().Header().Set("Content-Type", "application/json")
		// SUPA_ANON_KEY
		c.Response().Header().Set("apikey", os.Getenv("SUPA_ANON_KEY"))
		return handle(c)
	}
}

func NewHTTPHeaderHandler() *HTTPHeaderHandler {
	return &HTTPHeaderHandler{}
}
