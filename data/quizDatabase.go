package data

import (
	"github.com/hashicorp/go-hclog"
)

type QuizzesDB struct {
	loggger hclog.Logger
}

func NewQuizzesDB(logger hclog.Logger) *QuizzesDB {
	quizzesDB := &QuizzesDB{logger}
	return quizzesDB
}
