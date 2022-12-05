package handlers

import (
	"net/http"
	"quiz-app/quiz-api/data"
)

// GetAll handles GET requests and returns all current quizzes
func (quizHandler *QuizHandler) GetAllQuizResults(rw http.ResponseWriter, r *http.Request) {
	quizHandler.logger.Debug("Get all records")
	rw.Header().Add("Content-Type", "application/json")

	quizResults,err := quizHandler.quizResultModels.GetAllResults()

	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	err = data.ToJSON(quizResults, rw)
	if err != nil {
		// we should never be here but log the error just incase
		quizHandler.logger.Error("Unable to serializing quizzes", "error", err)
	}
}