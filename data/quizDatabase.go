package data

import (
	"fmt"

	"github.com/hashicorp/go-hclog"
)

type quiz struct {
	ID          int
	Name        string
	Description string
}

type question struct {
	ID              int
	Question        string
	correctChoiceID int
	quizID          int
}

type choice struct {
	ID         int
	Choice     string
	questionID int
}

type Answer struct {
	ID               int
	quizResultID     int
	questionID       int
	correctChoiceID  int
	selectedChoiceID int
	result           int //0=empty, 1=true, 2=false
}

type QuizResult struct {
	ID                  int
	quizID              int
	userID              int
	totalCorrectAnswers int
}

type user struct {
	ID       int
	userName string
}

type quizzes []*quiz
type questions []*question
type choices []*choice
type answers []*Answer
type QuizResults []*QuizResult
type users []*user

type QuizzesDB struct {
	loggger      hclog.Logger
}

func NewQuizzesDB(logger hclog.Logger) *QuizzesDB {
	quizzesDB := &QuizzesDB{logger}
	return quizzesDB
}

var quizzesList = quizzes{
	&quiz{ID: 1,
		Name:        "General Quiz",
		Description: "This is a general quiz. You can find questions from any topic."},
	&quiz{ID: 2,
		Name:        "Game Quiz",
		Description: "This quiz has questions that specific to video games."},
}

var questionsList = questions{
	&question{ID: 1,
		Question:        "What is the best game in the world?",
		correctChoiceID: 1,
		quizID:          1},
	&question{ID: 2,
		Question:        "What is the capital of Malta",
		correctChoiceID: 0,
		quizID:          1},
	&question{ID: 3,
		Question:        "What is the capital of Turkey",
		correctChoiceID: 0,
		quizID:          1},
	&question{ID: 4,
		Question:        "Test What is the best game in the world?",
		correctChoiceID: 1,
		quizID:          2},
	&question{ID: 5,
		Question:        "Test What is the capital of Malta",
		correctChoiceID: 1,
		quizID:          2},
	&question{ID: 6,
		Question:        "Test What is the capital of Turkey",
		correctChoiceID: 1,
		quizID:          2},
}

var choicesList = choices{
	&choice{ID: 1,
		Choice:     "God Of War",
		questionID: 1},
	&choice{ID: 2,
		Choice:     "GTA 5",
		questionID: 2},
	&choice{ID: 3,
		Choice:     "NFS",
		questionID: 2},
	&choice{ID: 4,
		Choice:     "Mario",
		questionID: 2},
	&choice{ID: 5,
		Choice:     "Call Of Duty",
		questionID: 1},
	&choice{ID: 5,
		Choice:     "Ankara",
		questionID: 3},
	&choice{ID: 1,
		Choice:     "God Of War",
		questionID: 4},
	&choice{ID: 2,
		Choice:     "GTA 5",
		questionID: 4},
	&choice{ID: 3,
		Choice:     "NFS",
		questionID: 5},
	&choice{ID: 4,
		Choice:     "Mario",
		questionID: 5},
	&choice{ID: 5,
		Choice:     "Call Of Duty",
		questionID: 6},
	&choice{ID: 5,
		Choice:     "Ankara",
		questionID: 6},
}

var answersList = answers{
	&Answer{ID: 1, quizResultID: 1,
		questionID: 1, correctChoiceID: 1, selectedChoiceID: 1},
	&Answer{ID: 2, quizResultID: 1,
		questionID: 2, correctChoiceID: 1, selectedChoiceID: 2},
	&Answer{ID: 3, quizResultID: 1,
		questionID: 2, correctChoiceID: 1, selectedChoiceID: 1},
	&Answer{ID: 1, quizResultID: 2,
		questionID: 1, correctChoiceID: 1, selectedChoiceID: 3},
	&Answer{ID: 2, quizResultID: 2,
		questionID: 2, correctChoiceID: 1, selectedChoiceID: 2},
	&Answer{ID: 3, quizResultID: 2,
		questionID: 2, correctChoiceID: 1, selectedChoiceID: 1},
}

var quizResultsList = QuizResults{
	&QuizResult{ID: 1, quizID: 1, userID: 1, totalCorrectAnswers: 2},
	&QuizResult{ID: 2, quizID: 1, userID: 1, totalCorrectAnswers: 1},
}

var usersList = users{
	&user{ID: 1,
		userName: "Emre"},
}

func GetAllQuizzes() quizzes {
	return quizzesList
}

func GetQuiz(id int) (*quiz, error) {

	for _, quiz := range quizzesList {
		if quiz.ID == id {
			return quiz, nil
		}
	}

	return nil, &ErrorQuizNotFound{id}
}

func GetAllQuestions() questions {
	return questionsList
}

func GetQuizQuestions(quizId int) (*questions, error) {

	quizQuestions := questions{}

	for _, question := range questionsList {
		if question.quizID == quizId {
			quizQuestions = append(quizQuestions, question)
		}
	}

	if len(quizQuestions) == 0 {
		return nil, &ErrorQuestionNotFound{quizId}
	}

	return &quizQuestions, nil
}

func GetAllChoices() choices {
	return choicesList
}

func GetQuestionChoices(questionId int) (*choices, error) {

	questionChoices := choices{}

	for _, choice := range choicesList {
		if choice.questionID == questionId {
			questionChoices = append(questionChoices, choice)
		}
	}

	if len(questionChoices) == 0 {
		return nil, &ErrorChoiceNotFound{questionId}
	}

	return &questionChoices, nil
}

func GetAllAnswers() answers {
	return answersList
}

func GetAnswer(answerId int) (*Answer, error) {

	for _, answer := range answersList {
		if answer.ID == answerId {
			return answer, nil
		}
	}

	return nil, &ErrorAnswerNotFound{answerId: answerId}
}

func GetAllUsers() users {
	return usersList
}

func (quiz quiz) String() string {
	return fmt.Sprintf("{%v %v}", quiz.ID, quiz.Name)
}

func (question question) String() string {
	return fmt.Sprintf("{%v %v %v %v}", question.ID, question.Question, question.correctChoiceID, question.quizID)
}

func (choice choice) String() string {
	return fmt.Sprintf("{%v %v %v}", choice.ID, choice.Choice, choice.questionID)
}

func (answer Answer) String() string {
	return fmt.Sprintf("{%v %v %v %v %v %v}", answer.ID, answer.quizResultID, answer.questionID, answer.correctChoiceID, answer.selectedChoiceID, answer.result)
}

func (user user) String() string {
	return fmt.Sprintf("{%v %v}", user.ID, user.userName)
}

type ErrorQuizNotFound struct {
	id int
}

func (err *ErrorQuizNotFound) Error() string {
	return fmt.Sprintf("Quiz id: %v could not found.", err.id)
}

type ErrorQuestionNotFound struct {
	quizId int
}

func (err *ErrorQuestionNotFound) Error() string {
	return fmt.Sprintf("Could not found any question. Quiz id : %v", err.quizId)
}

type ErrorChoiceNotFound struct {
	questionId int
}

func (err *ErrorChoiceNotFound) Error() string {
	return fmt.Sprintf("Could not found any choice. Question id : %v", err.questionId)
}

type ErrorUserNotFound struct {
	userId int
}

func (err *ErrorUserNotFound) Error() string {
	return fmt.Sprintf("Could not found user. User id : %v", err.userId)
}

type ErrorAnswerNotFound struct {
	answerId int
}

func (err *ErrorAnswerNotFound) Error() string {
	return fmt.Sprintf("Could not found answer. Answer id : %v", err.answerId)
}

type ErrorQuestionsNotFound struct {
	quizId int
}

func (err *ErrorQuestionsNotFound) Error() string {
	return fmt.Sprintf("Could not found questions for quiz. Quiz id : %v", err.quizId)
}

type ErrorQuizResultNotFound struct {
	quizResultID int
}

func (err *ErrorQuizResultNotFound) Error() string {
	return fmt.Sprintf("Could not found answered quiz. Quiz result id : %v", err.quizResultID)
}

type ErrorAnswersNotFound struct {
	quizResultID int
}

func (err *ErrorAnswersNotFound) Error() string {
	return fmt.Sprintf("Could not found answers for quiz result. Quiz result id : %v", err.quizResultID)
}

func GetMaxAnswersId() int {
	maxId := 0

	for _, answer := range answersList {
		if answer.ID > maxId {
			maxId = answer.ID
		}
	}

	return maxId
}

func GetMaxQuizResultId() int {
	maxId := 0

	for _, quizResult := range quizResultsList {
		if quizResult.ID > maxId {
			maxId = quizResult.ID
		}
	}

	return maxId
}

func CreateNewAnswer(quizResultID int, question *question, selectedChoiceID int) int {
	id := GetMaxAnswersId() + 1
	quizResult,err := GetQuizResultsByQuizResultID(quizResultID)

	if err!= nil{
		fmt.Println(err)
	}

	result := 0
	if selectedChoiceID == question.correctChoiceID {
		result = 1
		quizResult.UpdateTotalCorrectAnswer()
	} else if selectedChoiceID != question.correctChoiceID && selectedChoiceID != -1 {
		result = 2
	}

	newAnswer := Answer{ID: id,
		quizResultID:     quizResultID,
		questionID:       question.ID,
		correctChoiceID:  question.correctChoiceID,
		selectedChoiceID: selectedChoiceID,
		result:           result}
	answersList = append(answersList, &newAnswer)

	return id
}

func CreateNewQuizResult(quizID int, userID int) QuizResult {

	id := GetMaxQuizResultId() + 1
	newQuizResult := QuizResult{ID: id,
		quizID:              quizID,
		userID:              userID,
		totalCorrectAnswers: 0}
	quizResultsList = append(quizResultsList, &newQuizResult)

	return newQuizResult
}

func GetUser(userId int) (*user, error) {
	for _, user := range usersList {
		if userId == user.ID {
			return user, nil
		}
	}

	return nil, &ErrorUserNotFound{userId}
}

func (user *user) GetUserName() string {
	return user.userName
}

func (question *question) GetCorrectAnswer() int {
	return question.correctChoiceID
}

func GetCorrectChoiceByQuestionID(questionID int) (int, error) {

	for _, question := range questionsList {
		if question.ID == questionID {
			return question.correctChoiceID, nil
		}
	}

	return 0, &ErrorQuestionNotFound{questionID}
}

func GetQuizByID(quizID int) (*quiz, error) {

	for _, quiz := range quizzesList {
		if quiz.ID == quizID {
			return quiz, nil
		}
	}

	return nil, &ErrorQuizNotFound{quizID}
}

func GetQuestionByQuizID(quizID int) (*questions, error) {

	questions := questions{}
	for _, question := range questionsList {
		if question.quizID == quizID {
			questions = append(questions, question)
		}
	}

	if len(questions) != 0 {
		return &questions, nil
	}

	return nil, &ErrorQuestionsNotFound{quizID}
}

func GetQuizResultsByQuizResultID(quizResultID int) (*QuizResult, error) {

	for _, quizResult := range quizResultsList {
		if quizResult.ID == quizResultID {
			return quizResult, nil
		}
	}

	return nil, &ErrorQuizResultNotFound{quizResultID}
}

func (quizResult *QuizResult) GetQuizID() int {
	return quizResult.quizID
}

func (quizResult *QuizResult) GetUserID() int {
	return quizResult.userID
}

func (quizResult *QuizResult) GetTotalCorrectAnswers() int {
	return quizResult.totalCorrectAnswers
}

func (quizResult *QuizResult) GetAnswers() (*answers, error) {
	answers := answers{}
	for _, answer := range answersList {
		if answer.quizResultID == quizResult.ID {
			answers = append(answers, answer)
		}
	}

	if len(answers) == 0 {
		return nil, &ErrorAnswersNotFound{quizResult.ID}
	}

	return &answers, nil
}

func (answer *Answer) GetSelectedChoiceID() int {
	return answer.selectedChoiceID
}

func (answer *Answer) GetCorrectChoiceID() int {
	return answer.correctChoiceID
}

func (answer *Answer) GetResult() int {
	return answer.result
}

func GetAllQuizResults() *QuizResults {
	return &quizResultsList
}

func (quizResult *QuizResult) UpdateTotalCorrectAnswer() {
	fmt.Println("before quizResult.totalCorrectAnswers: ")
	fmt.Println(quizResult.totalCorrectAnswers)
	quizResult.totalCorrectAnswers++
	fmt.Println("after quizResult.totalCorrectAnswers: ")
	fmt.Println(quizResult.totalCorrectAnswers)
}
