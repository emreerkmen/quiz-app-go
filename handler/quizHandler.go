package handler

import (
	"quiz-app/quiz-api/data"

	"github.com/hashicorp/go-hclog"
)

// Quiz handler for geting and anwering Quizzes
type QuizHandler struct {
	logger         hclog.Logger
	validation     *data.Validation
	productDB      *data.QuizzesDB
}

// NewQuizHandler returns a new quiz handler with the given logger
func NewQuizHandler(logger hclog.Logger, validation *data.Validation, quizDB *data.QuizzesDB) *QuizHandler {
	return &QuizHandler{logger, validation, quizDB}
}