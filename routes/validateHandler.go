package routes

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/vukovlevi/kviz-go/db"
	"github.com/vukovlevi/kviz-go/views"
)

func parseIntValues(c echo.Context) (int, int, int, error) {
	id, err := strconv.Atoi(c.FormValue("game-id"))
	if err != nil {
		return -1, -1, -1, err
	}

	index, err := strconv.Atoi(c.Param("index"))
	if err != nil {
		return -1, -1, -1, err
	}

	answerIndex, err := strconv.Atoi(c.FormValue("answer-index"))
	if err != nil {
		return -1, -1, -1, err
	}

	return id, index, answerIndex, nil
}

func ValidateHandler(c echo.Context) error {
	id, index, answerIndex, err := parseIntValues(c)
	if err != nil {
		return render(c, views.Error("Szöveges kvíz id, kérdés sorszám vagy válasz sorszám"))
	}

	game, err := db.GetGameById(id)
	if err != nil {
		return render(c, views.Error("Nem létező kvíz id"))
	}

	rightAnser :=  (*game.Questions)[index].RightAnswerIndex == answerIndex
	nextIndex := index + 1
	if nextIndex >= game.Length {
		nextIndex = -1
	}
	return render(c, views.Answer(rightAnser, game.Id, nextIndex)) 
}