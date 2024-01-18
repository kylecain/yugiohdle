package model

import "gorm.io/gorm"

type Guess struct {
	gorm.Model
	ImageUrlCropped string
	Name            string
	Type            string
	FrameType       string
	Atk             int
	Def             int
	Level           int
	Race            string
	Attribute       string
}
