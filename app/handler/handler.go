package handler

import (
	"github.com/labstack/echo"
	
)

// AuthTest is AurtTest
func AuthTest() echo.HandlerFunc {
	return func(c echo.Context) error {

		return c.String(200, "")
	}
}
