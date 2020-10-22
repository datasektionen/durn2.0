package durn

import (
	"errors"
	"github.com/google/uuid"
	"sync"
)

var elections = make(map[uuid.UUID]*Election)
var mutex sync.Mutex

func CreateElection(name string, alternatives []Alternative) {
	mutex.Lock()
	defer mutex.Unlock()

	e := Election{
		Id:             uuid.New(),
		Name:           name,
		Alternatives:   alternatives,
		EligibleVoters: make(map[string]*EligibleVoter),
		Votes:          make(map[uuid.UUID]Vote),
	}

	elections[e.Id] = &e
}

func GetElections() []Election {
	mutex.Lock()
	defer mutex.Unlock()

	e := make([]Election, 0, len(elections))

	for _, election := range elections {
		e = append(e, *election)
	}

	return e
}

func AddEligibleVoters(electionId uuid.UUID, voterIds []string) {
	mutex.Lock()
	defer mutex.Unlock()

	election := elections[electionId]

	for _, id := range voterIds {
		if _, ok := election.EligibleVoters[id]; ok {
			continue
		}

		voter := EligibleVoter{
			Id:    id,
			Voted: false,
		}

		election.EligibleVoters[id] = &voter
	}

	elections[electionId] = election

}

func CastVote(electionId uuid.UUID, voterId string, alternative Alternative) error {
	mutex.Lock()
	defer mutex.Unlock()

	election, ok := elections[electionId]
	if !ok {
		return errors.New("election does not exist")
	}

	if !election.hasAlternative(alternative) {
		return errors.New("not valid alternative")
	}

	voter, ok := election.EligibleVoters[voterId]
	if !ok {
		return errors.New("voter does not exist")
	}
	if voter.Voted {
		return errors.New("voter already voted")
	}

	vote := Vote{
		Id:          uuid.New(),
		Alternative: alternative,
	}

	voter.Voted = true
	election.Votes[vote.Id] = vote

	return nil
}
