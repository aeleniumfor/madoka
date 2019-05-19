package handler

import (
	"github.com/labstack/echo"
	"github.com/madoka/app/common"
)

// AuthTest is AurtTest
func AuthTest() echo.HandlerFunc {
	return func(c echo.Context) error {

		return c.String(200, "")
	}
}

// CreateContainer is Create Container Request
func CreateContainer() echo.HandlerFunc {
	return func(c echo.Context) error {

		return c.String(200, "")
	}
}
