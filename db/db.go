package db

import (
	"fmt"

	"github.com/vukovlevi/kviz-go/models"
)

var questions []models.Question
var games []models.Game

func PopulateQuestions() {
    questions = models.PopulateQuestions()
}

func PopulateGames() {
    games = append(games, models.Game{Id: 1, Title: "A nagy kvíz", Topic: "Minden ami magyar", Length: len(questions), Questions: &questions})
}

func GetGames() *[]models.Game {
	return &games
}

func GetGameById(id int) (*models.Game, error) {
	for _, game := range games {
		if game.Id == id {
			return &game, nil
		}
	}

	return nil, fmt.Errorf("Game %d not found", id)
}

func AddQuestion(question models.Question) {
	questions = append(questions, question)
	games[0].Length++
}
