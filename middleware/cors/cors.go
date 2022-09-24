package cors

import (
	"github.com/labstack/echo"
)

type CORSHandler struct{}

func (m *CORSHandler) HandleCORS(handle echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("Access-Control-Allow-Origin", "*")
		return handle(c)
	}
}

func NewCORSHandler() *CORSHandler {
	return &CORSHandler{}
}
