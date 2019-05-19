package main

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/madoka/app/handler"
)

func main() {
	s := echo.New()
	s.Use(middleware.Logger())

	s.GET("/auth/test", handler.AuthTest())
	s.GET("/container/list/id",handler.CreateContainer())
	fmt.Println("madoka")
	s.Start(":8081")
}
