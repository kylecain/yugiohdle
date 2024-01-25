package repository

import (
	"regexp"
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
    re := regexp.MustCompile(`[^a-zA-Z0-9]`)
    filteredTerm := re.ReplaceAllString(searchTerm, "%")
	expression := "%" + filteredTerm + "%"

	var cards []model.Card
	r.DB.Where("name LIKE ? AND type LIKE '%monster%'", expression).
		Order("LENGTH(name)").
		Limit(10).
		Find(&cards)

	return cards
}

func (r *CardRepository) GetCurrentCard() string {
	var guess model.Guess
	r.DB.Last(&guess)
	return guess.Name
}

func (r *CardRepository) SearchCard(cardName string) model.Card {
	var card model.Card
	r.DB.Where("name = ?", cardName).
		First(&card)

	return card
}
