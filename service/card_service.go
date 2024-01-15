package service

import (
    "io"
    "net/http"
    "encoding/json"
    "yugiohdle/model"
)

type CardService struct {

}

func NewCardService() *CardService {
    return &CardService{}
}

func (s *CardService) GetRandomCard() (model.Card, error) {
	randomCard := "https://db.ygoprodeck.com/api/v7/randomcard.php"

	response, err := http.Get(randomCard)
	if err != nil {
		return model.Card{}, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return model.Card{}, err
	}

	var card model.Card

	err = json.Unmarshal(body, &card)
	if err != nil {
		return model.Card{}, err
	}

    return card, nil
}
