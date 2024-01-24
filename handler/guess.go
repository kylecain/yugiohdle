package handler

import (
	"fmt"
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
	fmt.Println(c)
	cardName := c.Request().FormValue("search")

	card := h.CardService.SearchCard(cardName)
	guess := model.Guess{
		ImageUrlCropped: "",
		Name:            card.Name,
		Type:            card.Type,
		FrameType:       card.FrameType,
		Atk:             card.Atk,
		Def:             card.Def,
		Level:           card.Level,
		Race:            card.Race,
		Attribute:       card.Attribute,
	}
	return render(c, component.GuessRow(guess))
}
