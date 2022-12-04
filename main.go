package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"quiz-app/quiz-api/data"
	"quiz-app/quiz-api/handlers"
	"quiz-app/quiz-api/model"
	"time"

	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
)

func main() {
	fmt.Println("Quiz app started.")

	//Create Server
	logger := hclog.Default()
	validation := data.NewValidation()

	// create quiz models instance
	quizModel := model.NewQuizzesModels(logger)
	quizResultModels := model.NewQuizResultModels(logger)
	makeCuopleOfQuizzes(*quizModel,*quizResultModels)

	// create the handlers
	quizHandler := handlers.NewQuizHandler(logger, validation, quizModel, quizResultModels)

	// create a new router and register the handlers
	serverMux := mux.NewRouter()

	// handlers for API
	getRouter := serverMux.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/quizzes", quizHandler.GetAllQuizzes)
	getRouter.HandleFunc("/quizResults", quizHandler.GetAllQuizResults)

	// CORS
	ch := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"*"}))

	// create a new server
	server := http.Server{
		Addr:         ":9090",                                               // configure the bind address
		Handler:      ch(serverMux),                                         // set the default handler
		ErrorLog:     logger.StandardLogger(&hclog.StandardLoggerOptions{}), // set the logger for the server
		ReadTimeout:  5 * time.Second,                                       // max time to read request from the client
		WriteTimeout: 10 * time.Second,                                      // max time to write response to the client
		IdleTimeout:  120 * time.Second,                                     // max time for connections using TCP Keep-Alive
	}

	// start the server
	go func() {
		logger.Info("Starting server on port 9090")

		if err := server.ListenAndServe(); err != nil {
			logger.Error("Error starting server", "error", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	// Block until a signal is received.
	sig := <-c
	logger.Debug("Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	if err := server.Shutdown(ctx); err != nil {
		logger.Error("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}

func makeCuopleOfQuizzes(quizModels model.QuizzesModels, quizResultModels model.QuizResultModels) {
	quizModels.GetAllQuizzes()
	//model.GetQuiz(1)
	fmt.Println("Answer a quiz.")
	quizResultID, err := model.AnswerQuiz(1, 1, []int{1, 1, -1})
	fmt.Printf("quizResultID : %v\n", quizResultID)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Get a quiz result")
	quizResultModels.GetResult(quizResultID)

	quizResultID, err = model.AnswerQuiz(2, 1, []int{-1, -1, -1})
	fmt.Printf("quizResultID : %v\n", quizResultID)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Get a quiz result")
	quizResultModels.GetResult(quizResultID)

	quizResultID, err = model.AnswerQuiz(1, 1, []int{1, 0, 0})
	fmt.Printf("quizResultID : %v\n", quizResultID)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Get a quiz result")
	quizResultModels.GetResult(quizResultID)

	quizResultID, err = model.AnswerQuiz(1, 1, []int{1, 1, 0})
	fmt.Printf("quizResultID : %v\n", quizResultID)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Get a quiz result")
	quizResultModels.GetResult(quizResultID)

	quizResultID, err = model.AnswerQuiz(1, 1, []int{1, 1, 0})
	fmt.Printf("quizResultID : %v\n", quizResultID)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Get a quiz result")
	quizResultModels.GetResult(quizResultID)

	quizResultID, err = model.AnswerQuiz(1, 1, []int{1, 1, 0})
	fmt.Printf("quizResultID : %v\n", quizResultID)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Get a quiz result")
	quizResultModels.GetResult(quizResultID)
}
