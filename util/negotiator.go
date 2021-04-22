package util

import (
	"github.com/labstack/echo/v4"
)

func Negotiate(c echo.Context, code int, i interface{}) error {
	mediaType := c.QueryParam("mediaType")

	switch mediaType {
	case "xml":
		return c.XML(code, i)
	case "json":
		return c.JSON(code, i)
	default:
		return c.JSON(code, i)
	}
}
