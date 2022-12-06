package data

import "fmt"

type choice struct {
	ID         int
	Choice     string
	questionID int
}

type choices []*choice

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
	&choice{ID: 6,
		Choice:     "Ankara",
		questionID: 3},
	&choice{ID: 7,
		Choice:     "İstanbul",
		questionID: 3},
	&choice{ID: 8,
		Choice:     "İzmir",
		questionID: 3},
	&choice{ID: 9,
		Choice:     "God Of War",
		questionID: 4},
	&choice{ID: 10,
		Choice:     "GTA 5",
		questionID: 4},
	&choice{ID: 11,
		Choice:     "NFS",
		questionID: 5},
	&choice{ID: 12,
		Choice:     "Mario",
		questionID: 5},
	&choice{ID: 13,
		Choice:     "Call Of Duty",
		questionID: 6},
	&choice{ID: 14,
		Choice:     "Ankara",
		questionID: 6},
}

func (choice choice) String() string {
	return fmt.Sprintf("{%v %v %v}", choice.ID, choice.Choice, choice.questionID)
}

type ErrorChoiceNotFound struct {
	questionId int
}

func (err *ErrorChoiceNotFound) Error() string {
	return fmt.Sprintf("Could not found any choice. Question id : %v", err.questionId)
}

func GetAllChoices() choices {
	return choicesList
}
