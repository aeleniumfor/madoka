package handler

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

// Net is controll network
type Net struct {
	ContainerID string `json:"container_id"`
}

// AuthTest is AurtTest
func AuthTest() echo.HandlerFunc {
	return func(c echo.Context) error { // なにか考える
		return c.String(http.StatusOK, "")
	}
}

// AddNetworkHandler is network Adding
func AddNetworkHandler() echo.HandlerFunc {
	return func(c echo.Context) error { // コンテナへIPを追加
		u := new(Net)
		if err := c.Bind(u); err != nil {
			return err
        }
        fmt.Println(u.ContainerID)
		return c.JSON(http.StatusOK, u)
	}
}

// DelNetworkHandler is network Adding
func DelNetworkHandler() echo.HandlerFunc {
	return func(c echo.Context) error { // コンテナへIPを追加
		u := new(Net)
		if err := c.Bind(u); err != nil {
			return err
        }
        fmt.Println(u.ContainerID)
		return c.JSON(http.StatusOK, u)
	}
}


// NetworkList is AurtTest
func NetworkList() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(200,"OK")
	}
}
