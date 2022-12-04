package handler

import (
	"quiz-app/quiz-api/data"

	"github.com/hashicorp/go-hclog"
)

// Quiz handler for geting and anwering Quizzes
type Products struct {
	logger         hclog.Logger
	validation     *data.Validation
	productDB      *data.
}