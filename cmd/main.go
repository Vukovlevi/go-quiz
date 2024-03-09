package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/vukovlevi/kviz-go/db"
	"github.com/vukovlevi/kviz-go/middleware"
	"github.com/vukovlevi/kviz-go/routes"
)

func main() {
	e := echo.New()

	e.GET("/", routes.IndexHandler)
	e.GET("/play", routes.HandlePlayIndex)
	e.GET("/play/:id", routes.HandlePlayById)
	e.GET("/play/:id/:index", routes.HandleQuestionByIndex)
	e.POST("/validate/:index", routes.ValidateHandler)

    g := e.Group("/new")
    g.Use(middleware.CheckAuth)
	g.GET("", routes.HandleNewQuestionIndex)
	g.POST("", routes.HandleNewQuestion)

    e.GET("/login", routes.HandleLoginIndex)
    e.POST("/login", routes.HandleLogin)

	e.Static("/", "public")

	db.PopulateQuestions()
	db.PopulateGames()
    db.CreateUser()

	log.Fatal(e.Start(":8080"))
}
