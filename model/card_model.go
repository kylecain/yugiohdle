package model

import "gorm.io/gorm"

type CardImage struct {
    gorm.Model
    CardID uint
    Id                int    `json:"id"`
    Image_url         string `json:"image_url"`
    Image_url_small   string `json:"image_url_small"`
    Image_url_cropped string `json:"image_url_cropped"`
}
