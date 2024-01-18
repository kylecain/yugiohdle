package repository

import (
	"fmt"
	"yugiohdle/model"

	"gorm.io/gorm"
)

type CardRepository struct {
    DB *gorm.DB
}

func NewCardRepository(db *gorm.DB) *CardRepository {
    return &CardRepository{
        DB: db,
    }
}

func (r *CardRepository) SearchCards(searchTerm string) []model.Card {
    var cards []model.Card
    expression := "%"+searchTerm+"%"
    fmt.Println(expression)
    r.DB.Where("name LIKE ?", expression).Find(&cards)

    return cards
}
