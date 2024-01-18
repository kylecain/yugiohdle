package handler

import (
	"yugiohdle/model"
	"yugiohdle/service"
	"yugiohdle/view/component"

	"github.com/labstack/echo/v4"
)

type GuessHandler struct {
	CardService service.CardService
}

func NewGuessHandler(cardService service.CardService) *GuessHandler {
	return &GuessHandler{
		CardService: cardService,
	}
}

func (h GuessHandler) HandleGuessShow(c echo.Context) error {
	return render(c, component.GuessResult([]model.Guess{}))
}
