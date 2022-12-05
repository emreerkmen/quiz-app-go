package model

import (
	"fmt"
	"quiz-app/quiz-api/data"

	"github.com/hashicorp/go-hclog"
)

//1, 1, []int{1, 1, 0}
type Answer struct {
	QuizID int
	UserID int
	SelectedChoices []int
}

type AnswerModels struct {
	logger hclog.Logger
}

func NewAnswerModels(logger hclog.Logger) *AnswerModels {
	answerModels := AnswerModels{logger: logger}
	return &answerModels
}

func(answerModel AnswerModels) AnswerQuiz(answer *Answer) (int, error) {
	quiz, err := data.GetQuizByID(answer.QuizID)

	if err != nil {
		return 0,&ErrorGeneric{err}
	}

	_, err = data.GetUser(answer.UserID)

	if err != nil {
		return 0,&ErrorGeneric{err}
	}

	questions, err := data.GetQuestionByQuizID(quiz.ID)

	if err != nil {
		return 0,&ErrorGeneric{err}
	}

	answerLength := len(answer.SelectedChoices)
	questionLength := len(*questions)

	if answerLength != questionLength {
		return 0, &ErrorAnswering{questionLength, answerLength}
	}

	quizResult := data.CreateNewQuizResult(answer.QuizID, answer.UserID)

	for index, selectedChoice := range answer.SelectedChoices {
		question := (*questions)[index]
		//choices length kontrolü yap
		choices, err := data.GetQuestionChoices(question.ID)
		if err != nil {
			return 0,&ErrorGeneric{err}
		}
		choicesLength := len(*choices)
		if selectedChoice+1 > choicesLength {
			return 0, &ErrorUnExpectedChoice{questionID: question.ID, choiceLenght: choicesLength, selectedChoice: selectedChoice}
		}
		data.CreateNewAnswer(quizResult.ID, question, selectedChoice)
	}

	return quizResult.ID, nil
}

type ErrorAnswering struct {
	questionLength int
	answerLength   int
}

func (err *ErrorAnswering) Error() string {
	return fmt.Sprintf("Answers and question lenght does not match. Question Length: %v, Answer Length: %v", err.questionLength, err.answerLength)
}

type ErrorUnExpectedChoice struct {
	questionID     int
	choiceLenght   int
	selectedChoice int
}

func (err *ErrorUnExpectedChoice) Error() string {
	return fmt.Sprintf("Selected a choise that does not eixst. Question ID: %v, Choice Length: %v, Selected Choice: %v", err.questionID, err.choiceLenght, err.selectedChoice+1)
}
