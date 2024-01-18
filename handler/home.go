package handler

import (
	"yugiohdle/service"
	"yugiohdle/view/home"

	"github.com/labstack/echo/v4"
)

type HomeHandler struct {
	CardService service.CardService
}

func NewHomeHandler(cardService service.CardService) *HomeHandler {
	return &HomeHandler{
		CardService: cardService,
	}
}

func (h HomeHandler) HandleHomeShow(c echo.Context) error {
	cardName := h.CardService.GetCurrentCard()

    return render(c, home.Show(cardName))
}
