package routes

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/vukovlevi/kviz-go/db"
	"github.com/vukovlevi/kviz-go/views"
)

func getNumberParam(c echo.Context, paramName string) (int, error) {
	return strconv.Atoi(c.Param(paramName))
}

func HandlePlayIndex(c echo.Context) error {
	games := db.GetGames()
	return render(c, views.Default(views.PlayIndex(games)))
}

func HandlePlayById(c echo.Context) error {
	id, err := getNumberParam(c, "id")
	if err != nil {
		return render(c, views.Default(views.Error("Érvénytelen kvíz id")))
	}

	game, err := db.GetGameById(id)
	if err != nil {
		return render(c, views.Default(views.Error("Kvíz az adott id-val nem található")))
	}

	return render(c, views.Default(views.PlayGame(game)))
}

func HandleQuestionByIndex(c echo.Context) error {
	id, err := getNumberParam(c, "id")
	if err != nil {
		return render(c, views.Error("Érvénytelen kvíz id"))
	}

	game, err := db.GetGameById(id)
	if err != nil {
		return render(c, views.Error("Kvíz az adott id-val nem található"))
	}

	index, err := getNumberParam(c, "index")
	if err != nil {
		return render(c, views.Error("Érvénytelen kérdés sorszám"))
	}

	if index >= game.Length || index < 0 {
		return render(c, views.Error("Nincs ilyen sorszámú kérdés"))
	}

	return render(c, views.Question(&(*game.Questions)[index], index))
}