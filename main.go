package main

import (
	"fmt"
	"quiz-app/quiz-api/model"
)

func main() {
	fmt.Println("Quiz app started.")
	makeCuopleOfQuizzes()
}


func makeCuopleOfQuizzes(){
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