package durn

import (
	"context"
	"durn2.0/auth"
	rl "durn2.0/requestLog"
	"durn2.0/util"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"sync"
)

var elections = make(map[uuid.UUID]*Election)
var mutex sync.Mutex

func CreateElection(ctx context.Context, name string, alternatives []Alternative) error {
	err := auth.IsAuthorized(ctx, "createElection")
	if err != nil {
		return err
	}

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

	rl.Info(ctx, fmt.Sprintf("Election created (name: \"%s\", id: %s)", e.Name, e.Id.String()))

	return nil
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

func GetEligibleVoters(_ context.Context, electionId uuid.UUID) ([]EligibleVoter, error) {
	mutex.Lock()
	defer mutex.Unlock()

	election, ok := elections[electionId]
	if !ok {
		return nil, errors.New("election does not exist")
	}

	voters := make([]EligibleVoter, 0, len(election.EligibleVoters))

	for _, voter := range election.EligibleVoters {
		voters = append(voters, *voter)
	}

	return voters, nil
}

func AddEligibleVoters(ctx context.Context, electionId uuid.UUID, voterIds []string) error {
	err := auth.IsAuthorized(ctx, "addEligibleVoters")
	if err != nil {
		return err
	}

	mutex.Lock()
	defer mutex.Unlock()

	election, ok := elections[electionId]
	if !ok {
		return errors.New("election does not exist")
	}

	added := 0
	alreadyAdded := 0

	for _, id := range voterIds {
		if _, ok := election.EligibleVoters[id]; ok {
			alreadyAdded += 1
			continue
		}

		voter := EligibleVoter{
			Id:    id,
			Voted: false,
		}

		election.EligibleVoters[id] = &voter

		added += 1
	}

	elections[electionId] = election

	rl.Info(ctx, fmt.Sprintf("Added %d voters (%d skippe due to already existing)", added, alreadyAdded))

	return nil
}

func CastVote(ctx context.Context, electionId uuid.UUID, alternative Alternative) error {
	mutex.Lock()
	defer mutex.Unlock()

	election, ok := elections[electionId]
	if !ok {
		return errors.New("election does not exist")
	}

	if !election.hasAlternative(alternative) {
		return errors.New("not valid alternative")
	}

	voterId := util.MustUser(ctx)
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
