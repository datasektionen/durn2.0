package durn

import (
	"time"

	"github.com/google/uuid"
)

type Election struct {
	id          uuid.UUID
	candidates  []uuid.UUID
	name        string
	isOpen      bool
	isFinalized bool
	openTime    time.Time
	closeTime   time.Time
}

type Candidate struct {
	id           uuid.UUID
	name         string
	presentation string
	elections    []uuid.UUID
}
