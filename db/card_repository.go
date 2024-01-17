package db

import (
	"context"
	"database/sql"
	"yugiohdle/model"
)

type CardRepository struct {
    db *sql.DB
}

func NewCardRepository(db *sql.DB) *CardRepository {
    return &CardRepository{db: db}
}

func (r *CardRepository) Create(ctx context.Context, card model.Card) error {
    statement := `

    `
    _, err := r.db.ExecContext(ctx)

    return err
}
