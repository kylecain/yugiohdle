package model

import "gorm.io/gorm"

type CardSet struct {
	gorm.Model
	CardID        uint
	SetName       string `json:"set_name"`
	SetCode       string `json:"set_code"`
	SetRarity     string `json:"set_rarity"`
	SetRarityCode string `json:"set_rarity_code"`
	SetPrice      string `json:"set_price"`
}
