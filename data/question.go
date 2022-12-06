package data

import "fmt"

type Question struct {
	ID              int
	Question        string
	correctChoiceID int
	quizID          int
}

type Questions []*Question

var questionsList = Questions{
	&Question{ID: 1,
		Question:        "What is the best game in the world?",
		correctChoiceID: 1,
		quizID:          1},
	&Question{ID: 2,
		Question:        "What is the capital of Malta",
		correctChoiceID: 0,
		quizID:          1},
	&Question{ID: 3,
		Question:        "What is the capital of Turkey",
		correctChoiceID: 0,
		quizID:          1},
	&Question{ID: 4,
		Question:        "Test What is the best game in the world?",
		correctChoiceID: 1,
		quizID:          2},
	&Question{ID: 5,
		Question:        "Test What is the capital of Malta",
		correctChoiceID: 1,
		quizID:          2},
	&Question{ID: 6,
		Question:        "Test What is the capital of Turkey",
		correctChoiceID: 1,
		quizID:          2},
}

type ErrorQuestionNotFound struct {
	quizId int
}

func (err *ErrorQuestionNotFound) Error() string {
	return fmt.Sprintf("Could not found any question. Quiz id : %v", err.quizId)
}

type ErrorQuestionsNotFound struct {
	quizId int
}

func (err *ErrorQuestionsNotFound) Error() string {
	return fmt.Sprintf("Could not found questions for quiz. Quiz id : %v", err.quizId)
}

func (question Question) String() string {
	return fmt.Sprintf("{%v %v %v %v}", question.ID, question.Question, question.correctChoiceID, question.quizID)
}

func GetAllQuestions() Questions {
	return questionsList
}

func GetQuestionChoices(questionId int) (*choices, error) {

	questionChoices := choices{}

	for _, choice := range choicesList {
		if choice.questionID == questionId {
			questionChoices = append(questionChoices, choice)
		}
	}

	if len(questionChoices) == 0 {
		return nil, &ErrorChoiceNotFound{questionId}
	}

	return &questionChoices, nil
}

func (question *Question) GetCorrectAnswer() int {
	return question.correctChoiceID
}
