package handler

import (
	"fmt"
	"io"
	"net/url"
	"strings"
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

func (h SearchHandler) HandleSearchShow(c echo.Context) error {
	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		fmt.Println(err)
	}

	rawSearchTerm := strings.Replace(string(body), "search=", "", -1)
	searchTerm, err := url.QueryUnescape(rawSearchTerm)

	if err != nil {
		fmt.Println(err)
	}

	cards := h.CardService.SearchCards(searchTerm)

	return render(c, component.SearchResult(cards))
}
