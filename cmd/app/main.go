package main

import (
	"yugiohdle/handler"
	"yugiohdle/repository"
	"yugiohdle/service"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)


func main() {
    app := echo.New()

	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

    cardRepository := repository.NewCardRepository(db)
    cardService := service.NewCardService(*cardRepository)
    homeHandler := handler.NewHomeHandler(*cardService)
    app.GET("/", homeHandler.HandleHomeShow)

    app.Start(":8080")
}
