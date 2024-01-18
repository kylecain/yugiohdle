package model

import "gorm.io/gorm"

type CardImage struct {
	gorm.Model
	CardID          uint
	Id              int    `json:"id"`
	ImageUrl        string `json:"image_url"`
	ImageUrlSmall   string `json:"image_url_small"`
	ImageUrlCropped string `json:"image_url_cropped"`
}
