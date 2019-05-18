package main

import (
	"github.com/labstack/echo/middleware"
	"github.com/labstack/echo"
	"github.com/madoka/app/handler"
	"fmt"
)

func main() {
	s := echo.New()
	s.Use(middleware.Logger())
	
	s.GET("/auth/test",handler.AuthTest())
	
	fmt.Println("madoka")
	s.Start(":8081")
}
