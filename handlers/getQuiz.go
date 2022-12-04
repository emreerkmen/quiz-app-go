package handlers

import "net/http"

// GetAll handles GET requests and returns all current quizzes
func (quizHandler *QuizHandler) GetAll(rw http.ResponseWriter, r *http.Request) {
	quizHandler.logger.Debug("Get all records")
	rw.Header().Add("Content-Type", "application/json")

	prods, err := quizHandler.quizzesDB
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	err = data.ToJSON(prods, rw)
	if err != nil {
		// we should never be here but log the error just incase
		p.l.Error("Unable to serializing product", "error", err)
	}
}