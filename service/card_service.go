package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"yugiohdle/model"
	"yugiohdle/repository"
)

type CardService struct {
    CardRepository repository.CardRepository
}

func NewCardService(cardRepository repository.CardRepository) *CardService {
    return &CardService{
        CardRepository: cardRepository,
    }
}

func (s *CardService) GetRandomCard() (model.Card, error) {
	url := "https://db.ygoprodeck.com/api/v7/randomcard.php"

	response, err := http.Get(url)
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

func (s *CardService) GetAllCards() ([]model.Card, error) {
    url := "https://db.ygoprodeck.com/api/v7/cardinfo.php"

    response, err := http.Get(url)

	if err != nil {
        fmt.Println(err)
		return []model.Card{}, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
        fmt.Println(err)
		return []model.Card{}, err
	}

    var data model.CardData

    err = json.Unmarshal(body, &data)
    if err != nil {
        fmt.Println(err)
        return []model.Card{}, err
    }
    
    return data.Data, nil
}
