package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/vukovlevi/kviz-go/db"
	"github.com/vukovlevi/kviz-go/routes"
)

func main() {
	e := echo.New()

	e.GET("/", routes.IndexHandler)
	e.GET("/play", routes.HandlePlayIndex)

	e.Static("/", "public")

	db.PopulateQuestions()
	db.PopulateGames()

	log.Fatal(e.Start(":8080"))
}