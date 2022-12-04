package model

import (
	"fmt"
	"quiz-app/quiz-api/data"

	"github.com/hashicorp/go-hclog"
)

type Quiz struct {
	ID          int
	Name        string
	Description string
	Questions   []*Question
}

type Question struct {
	Question string
	Choices  []*string
}

type Questions []*Question

type QuizzesModels struct {
	loggger hclog.Logger
}

func NewQuizzesModels(logger hclog.Logger) *QuizzesModels {
	quizzesModel := QuizzesModels{loggger: logger}
	return &quizzesModel
}

func (quizzesModels QuizzesModels) GetAllQuizzes() []Quiz {
	quizzesModel := []Quiz{}
	quizzes := data.GetAllQuizzes()

	for _, quiz := range quizzes {
		quizzesModel = append(quizzesModel, quizzesModels.GetQuiz(quiz.ID))
	}

	fmt.Println(quizzesModel)
	return quizzesModel
}

func (quizzesModels QuizzesModels) GetQuiz(quizId int) Quiz {
	quizModel := Quiz{}
	questionsModel := []*Question{}

	quiz, err := data.GetQuiz(quizId)

	if err != nil {
		fmt.Println(err)
	}

	quizModel.ID = quiz.ID
	quizModel.Name = quiz.Name
	quizModel.Description = quiz.Description

	questions, err := data.GetQuizQuestions(quizId)

	if err != nil {
		fmt.Println(err)
	}

	for _, question := range *questions {
		questionModel := Question{
			Question: question.Question,
			Choices:  GetChoicesStringArrays(question.ID),
		}
		questionsModel = append(questionsModel, &questionModel)
	}

	quizModel.Questions = questionsModel

	fmt.Println(quizModel)
	return quizModel
}

func GetChoicesStringArrays(questionId int) []*string {
	questionChoices := []*string{}

	choices, err := data.GetQuestionChoices(questionId)

	if err != nil {
		fmt.Println(err)
	}

	for _, choice := range *choices {
		questionChoices = append(questionChoices, &choice.Choice)
	}

	return questionChoices
}

func (quiz *Quiz) GetQuestions() []*Question {
	return quiz.Questions
}