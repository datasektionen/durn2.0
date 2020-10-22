package durn

import "github.com/google/uuid"

type Alternative string

type EligibleVoter struct {
	Id    string
	Voted bool
}

type Vote struct {
	Id          uuid.UUID
	Alternative Alternative
}

type Election struct {
	Id             uuid.UUID
	Name           string
	Alternatives   []Alternative
	EligibleVoters map[string]*EligibleVoter
	Votes          map[uuid.UUID]Vote
}

func (e *Election) hasAlternative(alternative Alternative) bool {
	for _, a := range e.Alternatives {
		if a == alternative {
			return true
		}
	}
	return false
}
