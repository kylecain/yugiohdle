package handler

import (
	"yugiohdle/service"
	"yugiohdle/view/home"

	"github.com/labstack/echo/v4"
)

type HomeHandler struct {

}

func(h HomeHandler) HandleHomeShow(c echo.Context) error {
    s := service.NewCardService()

    card, err := s.GetRandomCard()

    if err != nil {
        print(err)
    }

    return render(c, home.Show(card))
}
