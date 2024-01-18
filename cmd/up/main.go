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
	randomCard, _ := s.GetRandomCard()
	guess := model.Guess{
		ImageUrlCropped: randomCard.CardImages[0].ImageUrlCropped,
		Name:            randomCard.Name,
		Type:            randomCard.Type,
		FrameType:       randomCard.FrameType,
		Atk:             randomCard.Atk,
		Def:             randomCard.Def,
		Level:           randomCard.Level,
		Race:            randomCard.Race,
		Attribute:       randomCard.Attribute,
	}

	if err != nil {
		fmt.Println("error getting cards")
		return
	}

	db.AutoMigrate(
		&model.Card{},
		&model.CardSet{},
		&model.CardImage{},
		&model.CardPrice{},
		&model.Guess{},
	)

	result := db.CreateInBatches(&cards, 250)
	if result.Error != nil {
		panic("failed to save objects: " + result.Error.Error())
	}

	result = db.Create(&guess)
	if result.Error != nil {
		panic("failed to save objects: " + result.Error.Error())
	}
}
