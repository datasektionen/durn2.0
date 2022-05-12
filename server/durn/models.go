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

type ElectionState int

const (
	// The election is only visible to the election committee and auditors
	// Votes may not be cast
	Unpublished ElectionState = iota

	// Any authenticated user may see the election
	// Any authenticated user may cast a vote once
	//
	// An election in this state is considered published
	Voting

	// Any authenticated user may see the election
	// No new votes may be cast,
	// instead the result should be displayed to authenticated users
	//
	// An election in this state is considered published
	Closed
)

type Election struct {
	Id             uuid.UUID
	Name           string
	Alternatives   []Alternative
	EligibleVoters map[string]*EligibleVoter
	Votes          map[uuid.UUID]Vote
	State          ElectionState
}

func (e *Election) hasAlternative(alternative Alternative) bool {
	for _, a := range e.Alternatives {
		if a == alternative {
			return true
		}
	}
	return false
}
