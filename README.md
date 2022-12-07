# quiz-app-go
This is a simple quiz app that developed with golang.

It uses;
1. Gorilla Mux for Rest Api
2. Cobra for Cli

# App Start

1. Quiz Api
```bash
//in project root directory
$cd cmd/quizapi
$go run .
```

2. Quiz Cli
```bash
//in project root directory
$cd cmd/quizcli
$go run . --help
```

# End Points

## GET "/v1/quizzes"

Lists available quizzes

response:
```bash
[
    {
        "ID": 1,
        "name": "General Quiz",
        "description": "This is a general quiz. You can find questions from any topic.",
        "questions": [
            {
                "question": "What is the best game in the world?",
                "choices": [
                    "God Of War",
                    "Call Of Duty"
                ]
            },
            {
                "question": "What is the capital of Malta",
                "choices": [
                    "GTA 5",
                    "NFS",
                    "Mario"
                ]
            },
            {
                "question": "What is the capital of Turkey",
                "choices": [
                    "Ankara",
                    "İstanbul",
                    "İzmir"
                ]
            }
        ]
    },
    {
        "ID": 2,
        "name": "Game Quiz",
        "description": "This quiz has questions that specific to video games.",
        "questions": [
            {
                "question": "Test What is the best game in the world?",
                "choices": [
                    "God Of War",
                    "GTA 5"
                ]
            },
            {
                "question": "Test What is the capital of Malta",
                "choices": [
                    "NFS",
                    "Mario"
                ]
            },
            {
                "question": "Test What is the capital of Turkey",
                "choices": [
                    "Call Of Duty",
                    "Ankara"
                ]
            }
        ]
    }
]
```

## GET "/v1/quizResults"

Lists all quiz results

response:
```bash
[
    {
        "quizID": 1,
        "quizName": "General Quiz",
        "userName": "Emre",
        "questionAndAnswers": [
            {
                "question": "What is the best game in the world?",
                "selectedAnswer": "Call Of Duty",
                "correctAnswer": "Call Of Duty",
                "result": "Correct Answer :)"
            },
            {
                "question": "What is the capital of Malta",
                "selectedAnswer": "NFS",
                "correctAnswer": "GTA 5",
                "result": "Wrong Answer :("
            },
            {
                "question": "What is the capital of Turkey",
                "selectedAnswer": "İstanbul",
                "correctAnswer": "Ankara",
                "result": "Wrong Answer :("
            }
        ],
        "totalCorrectAnswers": 1,
        "status": 16
    },
    {
        "quizID": 1,
        "quizName": "General Quiz",
        "userName": "Emre",
        "questionAndAnswers": [
            {
                "question": "What is the best game in the world?",
                "selectedAnswer": "Call Of Duty",
                "correctAnswer": "Call Of Duty",
                "result": "Correct Answer :)"
            },
            {
                "question": "What is the capital of Malta",
                "selectedAnswer": "Mario",
                "correctAnswer": "NFS",
                "result": "Wrong Answer :("
            },
            {
                "question": "What is the capital of Turkey",
                "selectedAnswer": "İstanbul",
                "correctAnswer": "İstanbul",
                "result": "Correct Answer :)"
            }
        ],
        "totalCorrectAnswers": 1,
        "status": 16
    }
]    
```

## GET "/v1/quizResult/{id}"

lists the result with the given quiz result id

response:
```bash
{
    "quizID": 1,
    "quizName": "General Quiz",
    "userName": "Emre",
    "questionAndAnswers": [
        {
            "question": "What is the best game in the world?",
            "selectedAnswer": "Call Of Duty",
            "correctAnswer": "Call Of Duty",
            "result": "Correct Answer :)"
        },
        {
            "question": "What is the capital of Malta",
            "selectedAnswer": "NFS",
            "correctAnswer": "GTA 5",
            "result": "Wrong Answer :("
        },
        {
            "question": "What is the capital of Turkey",
            "selectedAnswer": "İstanbul",
            "correctAnswer": "Ankara",
            "result": "Wrong Answer :("
        }
    ],
    "totalCorrectAnswers": 1,
    "status": 16
}
```

## Post "/v1/answer"

Answer a quiz

request;
```bash
{
    "quizID": 1,
    "userId": 1,
    "selectedChoices": [
        0,
        0,
        0
    ]
}
```

response:
```bash
{
    "QuizResultID": 8
}
```
