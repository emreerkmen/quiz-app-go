package model

import (
	"fmt"
	"quiz-app/quiz-api/data"
)

func AnswerQuiz(quizId int, userId int, selectedChoices []int) (int, error) {
	quiz, err := data.GetQuizByID(quizId)

	if err != nil {
		fmt.Printf("Error when answering: %v", err)
	}

	fmt.Print("User id: ")
	fmt.Println(userId)
	_, err = data.GetUser(userId)

	if err != nil {
		fmt.Printf("Error when answering: %v", err)
	}

	questions, err := data.GetQuestionByQuizID(quiz.ID)

	if err != nil {
		fmt.Printf("Error when answering: %v", err)
	}

	answerLength := len(selectedChoices)
	questionLength := len(*questions)

	if answerLength != questionLength {
		return 0, &ErrorAnswering{questionLength, answerLength}
	}

	quizResult := data.CreateNewQuizResult(quizId, userId)

	for index, selectedChoice := range selectedChoices {
		question := (*questions)[index]
		//choices length kontrolü yap
		choices, err := data.GetQuestionChoices(question.ID)
		if err != nil {
			fmt.Println(err)
		}
		choicesLength := len(*choices)
		if selectedChoice+1 > choicesLength {
			return 0, &ErrorUnExpectedChoice{questionID: question.ID, choiceLenght: choicesLength, selectedChoice: selectedChoice}
		}
		data.CreateNewAnswer(quizResult.ID, question, selectedChoice)
	}

	answers := data.GetAllAnswers()
	fmt.Println(answers)

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
