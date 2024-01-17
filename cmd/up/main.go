package main

import (
	"fmt"
	"os"
	"yugiohdle/model"
	"yugiohdle/service"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	err := os.Remove("database.db")

	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	s := service.CardService{}

	cards, _ := s.GetAllCards()

	if err != nil {
		fmt.Println("error getting cards")
		return
	}

	db.AutoMigrate(&model.Card{})

	db.CreateInBatches(&cards, 250)
}
