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

func(h HomeHandler) HandleHomeShow(c echo.Context) error {
    card, err := h.CardService.GetRandomCard()

    if err != nil {
        print(err)
    }

    return render(c, home.Show(card))
}
