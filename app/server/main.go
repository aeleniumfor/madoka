package main

import (
	"github.com/labstack/echo/middleware"
	"github.com/labstack/echo"
	"fmt"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
    e.Use(middleware.Recover())

	fmt.Println("madoka")
	e.Start(":1323")
}
