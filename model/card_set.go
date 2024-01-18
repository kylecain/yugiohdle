package model

import "gorm.io/gorm"

type CardSet struct {
    gorm.Model
    CardID uint
    Set_name        string `json:"set_name"`
    Set_code        string `json:"set_code"`
    Set_rarity      string `json:"set_rarity"`
    Set_rarity_code string `json:"set_rarity_code"`
    Set_price       string `json:"set_price"`
}
