package handlers

import (
	"encoding/json"
	"net/http"
	"quiz-app/quiz-api/data"
	"quiz-app/quiz-api/model"
)

type QuizResult struct{
	QuizResultID int
}

// Answer a quiz
func (quizHandler *QuizHandler) AnswerQuiz(rw http.ResponseWriter, r *http.Request) {
	quizHandler.logger.Debug("Starting answering quiz")
	rw.Header().Add("Content-Type", "application/json")

	// Declare a new Answer struct.
	var answer model.Answer

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	err := json.NewDecoder(r.Body).Decode(&answer)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	quizHandler.logger.Debug("Post","Answer",answer)
	ID,err := quizHandler.answerModels.AnswerQuiz(&answer)
	quizResultID := QuizResult{QuizResultID: ID}

	if err!=nil{
		quizHandler.logger.Error("There are erros when answering quiz", "Answer", answer, "error", err)
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	err = data.ToJSON(quizResultID, rw)
	quizHandler.logger.Debug("","quizResultID", quizResultID)
	if err != nil {
		quizHandler.logger.Error("Unable to serializing quiz result", "error", err)
		return
	}

}