package main

import (
	"fmt"
	"net/http"
	"quiz-app/quiz-api/data"
	"quiz-app/quiz-api/handlers"
	"quiz-app/quiz-api/model"

	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
)

func main() {
	fmt.Println("Quiz app started.")
	makeCuopleOfQuizzes()

	//Create Server
	logger := hclog.Default()
	validation := data.NewValidation()

	// create database instance
	db := data.NewQuizzesDB(logger)

	// create the handlers
	quizHandler := handlers.NewQuizHandler(logger, validation, db)

	// create a new router and register the handlers
	serverMux := mux.NewRouter()

	// handlers for API
	getRouter := serverMux.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/quizzes", quizHandler.ListAll)
}

func makeCuopleOfQuizzes() {
	model.GetAllQuizzes()
	//model.GetQuiz(1)
	fmt.Println("Answer a quiz.")
	quizResultID, err := model.AnswerQuiz(1, 1, []int{1, 1, -1})
	fmt.Printf("quizResultID : %v\n", quizResultID)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Get a quiz result")
	model.GetResult(quizResultID)

	quizResultID, err = model.AnswerQuiz(2, 1, []int{-1, -1, -1})
	fmt.Printf("quizResultID : %v\n", quizResultID)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Get a quiz result")
	model.GetResult(quizResultID)

	quizResultID, err = model.AnswerQuiz(1, 1, []int{1, 0, 0})
	fmt.Printf("quizResultID : %v\n", quizResultID)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Get a quiz result")
	model.GetResult(quizResultID)

	quizResultID, err = model.AnswerQuiz(1, 1, []int{1, 1, 0})
	fmt.Printf("quizResultID : %v\n", quizResultID)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Get a quiz result")
	model.GetResult(quizResultID)

	quizResultID, err = model.AnswerQuiz(1, 1, []int{1, 1, 0})
	fmt.Printf("quizResultID : %v\n", quizResultID)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Get a quiz result")
	model.GetResult(quizResultID)

	quizResultID, err = model.AnswerQuiz(1, 1, []int{1, 1, 0})
	fmt.Printf("quizResultID : %v\n", quizResultID)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Get a quiz result")
	model.GetResult(quizResultID)
}
