package routes

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/vukovlevi/kviz-go/db"
	"github.com/vukovlevi/kviz-go/views"
)

func HandlePlayIndex(c echo.Context) error {
	games := db.GetGames()
	return render(c, views.PlayIndex(games))
}

func handlePlayById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return render(c, views.Default(views.Error("Érvénytelen kvíz id")))
	}

	game, err := db.GetGameById(id)
	if err != nil {
		return render(c, views.Default(views.Error("Kvíz az adott id-val nem található")))
	}
	//TODO: megjeleníteni egy adott game-et

}