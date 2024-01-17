package handler

import (
	"yugiohdle/service"
	"yugiohdle/view/component"

	"github.com/labstack/echo/v4"
)

type SearchHandler struct {
    CardService service.CardService
}

func NewSearchHandler(cardService service.CardService) *SearchHandler {
    return &SearchHandler{
        CardService: cardService,
    }
}

func(h SearchHandler) HandleSearchShow(c echo.Context) error {
    card, err := h.CardService.GetRandomCard()

    if err != nil {
        print(err)
    }

    return render(c, component.Show(card))
}
