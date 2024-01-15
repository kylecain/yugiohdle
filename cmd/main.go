package main

import (
	"yugiohdle/handler"

	"github.com/labstack/echo/v4"
)

type DB struct {}

func main() {
    app := echo.New()

    homeHandler := handler.HomeHandler{}
    app.GET("/", homeHandler.HandleHomeShow)

    app.Start(":8080")
}
