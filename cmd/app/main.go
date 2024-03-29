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

	searchHandler := handler.NewSearchHandler(*cardService)
	app.POST("/search", searchHandler.HandleSearchShow)

	guessHandler := handler.NewGuessHandler(*cardService)
	app.POST("/guess", guessHandler.HandleGuessShow)

	app.Static("/dist", "dist")
	app.Start(":8080")
}
