package model

import "gorm.io/gorm"

type Card struct {
	gorm.Model
	Id         int         `json:"id"`
	Name       string      `json:"name"`
	Type       string      `json:"type"`
	FrameType  string      `json:"frameType"`
	Desc       string      `json:"desc"`
	Atk        int         `json:"atk"`
	Def        int         `json:"def"`
	Level      int         `json:"level"`
	Race       string      `json:"race"`
	Attribute  string      `json:"attribute"`
	CardSets   []CardSet   `json:"card_sets"`
	CardImages []CardImage `json:"card_images"`
	CardPrices []CardPrice `json:"card_prices"`
}

type CardData struct {
	Data []Card `json:"data"`
}

