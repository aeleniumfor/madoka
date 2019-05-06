package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

// Net is controll network
type Net struct {
	Command      string `json:"Command"`
	ContainerID  string `json:"Container_Id"`
	NetNamespace string `json:"NetNs"`
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
		return c.String(200, "OK")
	}
}
