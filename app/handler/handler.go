package handler

import (
    "net/http"
    "github.com/labstack/echo"
)


// MainPage is 
func MainPage() echo.HandlerFunc {
    return func(c echo.Context) error {     //c をいじって Request, Responseを色々する 
        return c.String(http.StatusOK, "Hello World")
    }
}