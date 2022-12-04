package handlers

import (
	"quiz-app/quiz-api/data"
	"quiz-app/quiz-api/model"

	"github.com/hashicorp/go-hclog"
)

// Quiz handler for geting and anwering Quizzes
type QuizHandler struct {
	logger           hclog.Logger
	validation       *data.Validation
	quizzesModels    *model.QuizzesModels
	quizResultModels *model.QuizResultModels
}

// NewQuizHandler returns a new quiz handler with the given logger
func NewQuizHandler(logger hclog.Logger, validation *data.Validation, quizzesModels *model.QuizzesModels, quizResultModels *model.QuizResultModels) *QuizHandler {
	return &QuizHandler{logger, validation, quizzesModels, quizResultModels}
}
