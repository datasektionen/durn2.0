package durn

import (
	"context"
	"durn2.0/auth"
	rl "durn2.0/requestLog"
	"durn2.0/util"
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
		State:          Unpublished,
	}

	elections[e.Id] = &e

	rl.Info(ctx, fmt.Sprintf("Election created (name: \"%s\", id: %s)", e.Name, e.Id.String()))

	return nil
}

func GetUnpublishedElectionIds(ctx context.Context) ([]uuid.UUID, error) {
	err := auth.IsAuthorized(ctx, "getUnpublishedElectionIds")
	if err != nil {
		return nil, err
	}

	mutex.Lock()
	defer mutex.Unlock()

	ids := make([]uuid.UUID, 0, len(elections))

	for _, election := range elections {
		if election.State == Unpublished {
			ids = append(ids, election.Id)
		}
	}

	return ids, nil
}

func GetVotingElectionIds(ctx context.Context) ([]uuid.UUID, error) {
	// Any authenticated user is authorized
	if !auth.IsAuthenticated(ctx) {
		return nil, util.AuthenticationError("user is not authenticated")
	}

	mutex.Lock()
	defer mutex.Unlock()

	ids := make([]uuid.UUID, 0, len(elections))

	for _, election := range elections {
		if election.State == Voting {
			ids = append(ids, election.Id)
		}
	}

	return ids, nil
}

func GetClosedElectionIds(ctx context.Context) ([]uuid.UUID, error) {
	// Any authenticated user is authorized
	if !auth.IsAuthenticated(ctx) {
		return nil, util.AuthenticationError("user is not authenticated")
	}

	mutex.Lock()
	defer mutex.Unlock()

	ids := make([]uuid.UUID, 0, len(elections))

	for _, election := range elections {
		if election.State == Closed {
			ids = append(ids, election.Id)
		}
	}

	return ids, nil
}

func PublishElection(ctx context.Context, electionId uuid.UUID) error {
	err := auth.IsAuthorized(ctx, "publishElection")
	if err != nil {
		return err
	}

	mutex.Lock()
	defer mutex.Unlock()

	election, ok := elections[electionId]
	if !ok {
		return util.BadRequestError("election does not exist")
	}

	if election.State != Unpublished {
		return util.ConflictError("elections is not unpublished")
	}

	election.State = Voting

	return nil
}

func CloseElection(ctx context.Context, electionId uuid.UUID) error {
	err := auth.IsAuthorized(ctx, "closeElection")
	if err != nil {
		return err
	}

	mutex.Lock()
	defer mutex.Unlock()

	election, ok := elections[electionId]
	if !ok {
		return util.BadRequestError("election does not exist")
	}

	if election.State != Voting {
		return util.ConflictError("elections can not be closed as it is not in voting state")
	}

	election.State = Closed

	return nil
}

func GetEligibleVoters(ctx context.Context, electionId uuid.UUID) ([]EligibleVoter, error) {
	err := auth.IsAuthorized(ctx, "getEligibleVoters")
	if err != nil {
		return nil, err
	}

	mutex.Lock()
	defer mutex.Unlock()

	election, ok := elections[electionId]
	if !ok {
		return nil, util.BadRequestError("election does not exist")
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
		return util.BadRequestError("election does not exist")
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

func CastVote(ctx context.Context, electionId uuid.UUID, alternative Alternative) (*uuid.UUID, error) {
	// Any authenticated user is authorized
	if !auth.IsAuthenticated(ctx) {
		return nil, util.AuthenticationError("user is not authenticated")
	}

	mutex.Lock()
	defer mutex.Unlock()

	election, ok := elections[electionId]
	if !ok {
		return nil, util.BadRequestError("election does not exist")
	}

	if election.State != Voting {
		return nil, util.ConflictError("election is not open for voting right now")
	}

	if !election.hasAlternative(alternative) {
		return nil, util.BadRequestError("not valid alternative")
	}

	voterId := util.MustUser(ctx)
	voter, ok := election.EligibleVoters[voterId]
	if !ok {
		return nil, util.BadRequestError("voter does not exist")
	}
	if voter.Voted {
		return nil, util.ConflictError("voter already voted")
	}

	vote := Vote{
		Id:          uuid.New(),
		Alternative: alternative,
	}

	voter.Voted = true
	election.Votes[vote.Id] = vote

	return &vote.Id, nil
}

func GetVotes(ctx context.Context, electionId uuid.UUID) ([]Vote, error) {
	// Any authenticated user is authorized
	if !auth.IsAuthenticated(ctx) {
		return nil, util.AuthenticationError("user is not authenticated")
	}

	election, ok := elections[electionId]
	if !ok {
		return nil, util.BadRequestError("election does not exist")
	}

	if election.State != Closed {
		return nil, util.ConflictError("votes are only available when the election is closed")
	}

	votes := make([]Vote, 0, len(election.Votes))

	for _, vote := range election.Votes {
		votes = append(votes, vote)
	}

	return votes, nil
}
