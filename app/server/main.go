package main

import (
	"github.com/labstack/echo/middleware"
	"github.com/labstack/echo"
	"github.com/madoka/app/handler"
	"fmt"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/auth/test",handler.AuthTest())
	e.POST("/container/network",handler.AddNetworkHandler())

	fmt.Println("madoka")
	e.Start(":8081")
}
