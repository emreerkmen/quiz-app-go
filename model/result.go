package model

import (
	"fmt"
	"quiz-app/quiz-api/data"

	"github.com/hashicorp/go-hclog"
)

type Result struct {
	QuizId              int
	QuizName            string
	UserName            string
	QuestionAndAnswers  []*QuestionAndAnswer
	TotalCorrectAnswers int
	Status              int
}

type QuestionAndAnswer struct {
	Question       string
	SelectedAnswer string
	CorrectAnswer  string
	Result         string
}

type QuestionAndAnswers []*QuestionAndAnswer

type QuizResultModels struct {
	loggger hclog.Logger
}

func NewQuizResultModels(logger hclog.Logger) *QuizResultModels {
	quizResultModels := QuizResultModels{loggger: logger}
	return &quizResultModels
}

func (quizResultModesl QuizResultModels) GetAllResults() []Result {
	resultModels := []Result{}
	results := data.GetAllQuizResults()

	for _, result := range *results {
		resultModels = append(resultModels, quizResultModesl.GetResult(result.ID))
	}

	fmt.Println(resultModels)
	return resultModels
}

func(quizResultModesl QuizResultModels) GetResult(quizResultID int) Result {
	result := Result{}

	quizResult, err := data.GetQuizResultsByQuizResultID(quizResultID)
	fmt.Print("Quiz Result:  ")
	fmt.Println(quizResult)
	if err != nil {
		fmt.Println(err)
	}

	result.QuizId = quizResult.GetQuizID()
	quiz, err := data.GetQuizByID(result.QuizId)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("3")
	result.QuizName = quiz.Name
	fmt.Print("quizResult.GetUserID")
	fmt.Println(quizResult.GetUserID())
	user, err := data.GetUser(quizResult.GetUserID())
	fmt.Print("dönen user")
	fmt.Println(user)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("4")
	result.UserName = user.GetUserName()

	questions, err := data.GetQuizQuestions(result.QuizId)

	if err != nil {
		fmt.Println(err)
	}

	answers, err := quizResult.GetAnswers()

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("5")
	questionAndAnswers := QuestionAndAnswers{}
	for index, question := range *questions {
		questionText := question.Question
		choices, err := data.GetQuestionChoices(question.ID)
		fmt.Printf("Question id: %v\n", question.ID)
		fmt.Printf("Answer id: %v\n", (*answers)[index].ID)
		if err != nil {
			fmt.Println(err)
		}

		correctChoiceID := (*answers)[index].GetCorrectChoiceID()
		selectedChoiceID := (*answers)[index].GetSelectedChoiceID()

		fmt.Printf("correctChoiceID : %v\n", correctChoiceID)
		fmt.Printf("selectedChoiceID : %v\n", selectedChoiceID)
		fmt.Printf("Choices: %v\n", (*choices))

		var selectedAnswer string
		if selectedChoiceID != -1 {
			selectedAnswer = (*choices)[selectedChoiceID].Choice
		}
		correctAnswer := (*choices)[correctChoiceID].Choice

		questionAndAnswer := QuestionAndAnswer{Question: questionText,
			SelectedAnswer: selectedAnswer,
			CorrectAnswer:  correctAnswer}

		if selectedChoiceID == -1 {
			questionAndAnswer.Result = "Empty."
		} else if selectedChoiceID == correctChoiceID {
			questionAndAnswer.Result = "Correct Answer :)"
		} else {
			questionAndAnswer.Result = "Wrong Answer :("
		}

		questionAndAnswers = append(questionAndAnswers, &questionAndAnswer)
	}

	result.QuestionAndAnswers = questionAndAnswers
	result.TotalCorrectAnswers = quizResult.GetTotalCorrectAnswers()
	fmt.Print("quizResult.GetTotalCorrectAnswers() ")
	fmt.Printf("%v\n",quizResult.GetTotalCorrectAnswers())
	result.Status = CalculateStatus(quizResult.GetTotalCorrectAnswers(), quizResultID)

	fmt.Println(result)
	return result

}

func (queAndAns QuestionAndAnswer) String() string {
	return fmt.Sprintf("{%v %v %v %v }", queAndAns.Question, queAndAns.SelectedAnswer, queAndAns.CorrectAnswer, queAndAns.Result)
}

func CalculateStatus(currentTotalCorrectAnswer int, quizResultID int) int {
	fmt.Print("(currentTotalCorrectAnswer ")
	fmt.Printf("%v\n", currentTotalCorrectAnswer)
	quizResults := data.GetAllQuizResults()
	worstQuizResultsAmount := 0.0

	for _, quizResult := range *quizResults {
		fmt.Print("(GetTotalCorrectAnswers ")
		fmt.Printf("%v\n", quizResult.GetTotalCorrectAnswers())
		if quizResult.GetTotalCorrectAnswers() < currentTotalCorrectAnswer && quizResultID!=quizResult.ID {
			worstQuizResultsAmount++
		}
	}

	// When calculating do not count current quiz result
	fmt.Print("(worstquizResultsAmount ")
	fmt.Printf("%v\n", worstQuizResultsAmount)
	fmt.Print("(len(*quizResults)-1) ")
	fmt.Printf("%v\n", len(*quizResults)-1)

	if worstQuizResultsAmount == 0 {
		return 0
	}

	return int(worstQuizResultsAmount/float64((len(*quizResults) - 1)) * 100)
}
