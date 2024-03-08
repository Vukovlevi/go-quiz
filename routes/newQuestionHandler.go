package routes

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/vukovlevi/kviz-go/db"
	"github.com/vukovlevi/kviz-go/models"
	"github.com/vukovlevi/kviz-go/views"
)

func HandleNewQuestionIndex(c echo.Context) error {
	return render(c, views.Default(views.NewQuestionIndex()))
}

func HandleNewQuestion(c echo.Context) error {
	question := c.FormValue("question")
	answer1 := c.FormValue("answer-1")
	answer2 := c.FormValue("answer-2")
	answer3 := c.FormValue("answer-3")
	answer4 := c.FormValue("answer-4")
	rightAnswerIndex, err := strconv.Atoi(c.FormValue("right-answer"))
	if err != nil || rightAnswerIndex > 3 || rightAnswerIndex < 0 {
		return render(c, views.Error("Érvénytelen válasz index"))
	}

	if question == "" || answer1 == "" || answer2 == "" || answer3 == "" || answer4 == "" {
		return render(c, views.Error("Hiányzó paraméter"))
	}

	db.AddQuestion(models.Question{Question: question, Answers: []string{answer1, answer2, answer3, answer4}, RightAnswerIndex: rightAnswerIndex})

	return render(c, views.AddedQuestion())
}
