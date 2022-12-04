package model

import (
	"fmt"
	"quiz-app/quiz-api/data"
)

type Quiz struct {
	ID int
	Name  string
	Description string
	questions []Question
}

type Question struct {
	question string
	choices  []string
}

type Questions []Question

func GetAllQuizzes() ([]Quiz){
	quizzesModel := []Quiz{}
	quizzes := data.GetAllQuizzes()

	for _,quiz := range quizzes {
		quizzesModel = append(quizzesModel, GetQuiz(quiz.ID))
	}

	fmt.Println(quizzesModel)
	return quizzesModel
}

func GetQuiz(quizId int) (Quiz) {
	quizModel := Quiz{}
	questionsModel := []Question{}

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
           question: question.Question,
		   choices: GetChoicesStringArrays(question.ID),
		}
		questionsModel = append(questionsModel, questionModel)
	}

	quizModel.questions = questionsModel

	fmt.Println(quizModel)
	return quizModel
}

func GetChoicesStringArrays(questionId int) []string{
	questionChoices := []string{}

	choices,err := data.GetQuestionChoices(questionId)

	if err!=nil {
		fmt.Println(err)
	}

	for _, choice := range *choices {
		questionChoices  = append(questionChoices, choice.Choice)
	}

	return questionChoices
}

func(quiz *Quiz) GetQuestions() Questions{
	return quiz.questions
}

