package repository

import "gorm.io/gorm"

type CardRepository struct {
    DB *gorm.DB
}

func NewCardRepository(db *gorm.DB) *CardRepository {
    return &CardRepository{
        DB: db,
    }
}
