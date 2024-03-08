package models

type Question struct {
	Question string
	Answers []string
	RightAnswerIndex int
}

type Game struct {
	Id int
	Title string
	Topic string
	Length int
	Questions *[]Question
}

type User struct {
    Username string
    Password string
}


func PopulateQuestions() []Question {
    var questions []Question
	q1 := Question{"Mennyi 2+2?", []string{"1", "2", "3", "4"}, 3}
	questions = append(questions, q1)
	q2 := Question{"Mi a legnagyobb magyar trauma?", []string{"Mohács", "Trianon", "Gáspár Győző", "Gyúrcsány"}, 1}
	questions = append(questions, q2)
	q3 := Question{"Láttál-e már valaha csipkebokor rózsát?", []string{"Igen", "Nem", "Csipkebokor rózsa közt két szál majorannát", "Mi az a csipkebokor rózsa?"}, 2}
	questions = append(questions, q3)
	q4 := Question{"11, 12, 13...", []string{"Az ipari a legfaszább a világon", "14", "???", "14, 15, 16"}, 0}
	questions = append(questions, q4)
	q5 := Question{"Mi a vektor?", []string{"A mátrix titka", "Irányítot szakasz", "Viktor", "Orbán"}, 1}
	questions = append(questions, q5)
    return questions
}
