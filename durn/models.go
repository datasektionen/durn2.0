package durn

import (
	"time"

	"github.com/google/uuid"
)

type Election struct {
	Id          uuid.UUID
	Candidates  []uuid.UUID
	Name        string
	IsOpen      bool
	IsFinalized bool
	OpenTime    time.Time
	CloseTime   time.Time
}

type Candidate struct {
	Id           uuid.UUID
	Name         string
	Presentation string
	Elections    []uuid.UUID
}

type Vote struct {
	Hash     string
	Election uuid.UUID
	Ranking  uuid.UUID
	voteTime time.Time
}

type Voter string
