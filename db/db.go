package db

import (
	"fmt"
	"log"

	"github.com/vukovlevi/kviz-go/models"
	"golang.org/x/crypto/bcrypt"
)

var questions []models.Question
var games []models.Game
var user models.User

func PopulateQuestions() {
    questions = models.PopulateQuestions()
}

func PopulateGames() {
    games = append(games, models.Game{Id: 1, Title: "A nagy kvíz", Topic: "Minden ami magyar", Length: len(questions), Questions: &questions})
}

func CreateUser() {
    password, err := bcrypt.GenerateFromPassword([]byte("123"), 12)
    if err != nil {
        log.Fatal("Cannot create password")
    }
    user = models.User{Username: "vukovlevi", Password: string(password)}
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

func GetUserByUsername(username string) *models.User {
    if user.Username != username {
        return nil
    }
    return &user
}
