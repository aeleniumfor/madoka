package handler

import (
    "net/http"
    "github.com/labstack/echo"
)


// AuthTest is AurtTest
func AuthTest() echo.HandlerFunc {
    return func(c echo.Context) error { // なにか考える
        return c.String(http.StatusOK, "")
    }
}

// AddNetwork is network Adding
func AddNetwork() echo.HandlerFunc {
	return func(c echo.Context) error { // なにか考える
        return c.String(http.StatusOK, "")
    }
}