package quizcli

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"quiz-app/quiz-api/data"

	//"quiz-app/quiz-api/data"
	"strconv"
)

type Quiz struct {
	ID          int       `json:"ID"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Questions   Questions `json:"questions"`
}

type Question struct {
	Question string   `json:"question"`
	Choices  []string `json:"choices"`
}

type Quizzes []Quiz
type Questions []Question

type Result struct {
	QuizId              int                 `json:"quizID"`
	QuizName            string              `json:"quizName"`
	UserName            string              `json:"userName"`
	QuestionAndAnswers  []QuestionAndAnswer `json:"questionAndAnswers"`
	TotalCorrectAnswers int                 `json:"totalCorrectAnswers"`
	Status              int                 `json:"status"`
}

type QuestionAndAnswer struct {
	Question       string `json:"question"`
	SelectedAnswer string `json:"selectedAnswer"`
	CorrectAnswer  string `json:"correctAnswer"`
	Result         string `json:"result"`
}

type Answer struct {
	QuizID          int   `json:"quizID" validate:"required"`
	UserID          int   `json:"userID" validate:"required"`
	SelectedChoices []int `json:"selectedChoices" validate:"required,dive,min=-1"`
}

type QuestionAndAnswers []QuestionAndAnswer

type QuizResult struct{
	QuizResultID int
}

func Reverse(input string) (result string) {
	for _, c := range input {
		result = string(c) + result
	}
	return result
}

func Inspect(input string, digits bool) (count int, kind string) {
	if !digits {
		return len(input), "char"
	}
	return inspectNumbers(input), "digit"
}

func inspectNumbers(input string) (count int) {
	for _, c := range input {
		_, err := strconv.Atoi(string(c))
		if err == nil {
			count++
		}
	}
	return count
}

// Business logic
func GetQuizzes(beautify bool) (result string) {

	domain := "http://localhost:9090/v1/"
	uri := "quizzes"
	url := domain + uri

	resp, getErr := http.Get(url)
	if getErr != nil {
		log.Fatal(getErr)
	}

	quizzes := Quizzes{}

	return GetBeautifyJson(quizzes, resp.Body, beautify)
}

// Business logic
func GetQuizResults(beautify bool) (result string) {

	domain := "http://localhost:9090/v1/"
	uri := "quizResults"
	url := domain + uri

	resp, getErr := http.Get(url)
	if getErr != nil {
		log.Fatal(getErr)
	}

	quizzes := Quizzes{}

	return GetBeautifyJson(quizzes, resp.Body, beautify)
}

// Business logic
func PostAnswer(quizID string, userID string, selectedAnswers string, beautify bool) (result string) {

	intQuizID, err := strconv.Atoi(quizID)
	if err != nil {
		log.Fatal(err)
	}

	intUserID, err := strconv.Atoi(userID)
	if err != nil {
		log.Fatal(err)
	}

	intSelectedAnswers := []int{}

	var ans string
	for _, char := range selectedAnswers+"," {

		if char == ',' {

			intAns, err := strconv.Atoi(ans)
			if err != nil {
				log.Fatal(err)
			}

			intSelectedAnswers = append(intSelectedAnswers, intAns)
			ans = ""
			continue
		}
		ans = ans + string(char)
	}

	domain := "http://localhost:9090/v1/"
	uri := "answer"
	url := domain + uri

	answer := Answer{QuizID: intQuizID, UserID: intUserID, SelectedChoices: intSelectedAnswers}
	jsonAnswer, err := json.Marshal(answer)
	if err != nil {
		log.Fatal(err)
	}

	requestBody := bytes.NewBuffer(jsonAnswer)
	resp, err := http.Post(url, "application/json", requestBody)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	quizResultID:=QuizResult{}

	return GetBeautifyJson(quizResultID, resp.Body, beautify)
}

// Response body to string with beautify json
func GetBeautifyJson(i interface{}, r io.Reader, beautify bool) string {

	err := data.FromJSON(&i, r)
	if err != nil {
		log.Fatal("Deserializing quizzes", "error", err)
		return "Deserializing quizzes error"
	}

	if !beautify {
		return fmt.Sprintf("%v", i)
	}

	quizzesBea, err := json.MarshalIndent(i, "", " ")
	if err != nil {
		return ""
	}

	return string(quizzesBea)
}
