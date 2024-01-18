package model

import "gorm.io/gorm"

type CardPrice struct {
    gorm.Model
    CardID uint
    Cardmarket_price   string `json:"cardmarket_price"`
    Tcgplayer_price    string `json:"tcgplayer_price"`
    Ebay_price         string `json:"ebay_price"`
    Amazon_price       string `json:"amazon_price"`
    Coolstuffinc_price string `json:"coolstuffinc_price"`
}
