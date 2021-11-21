package handler

import (
	"net/http"
	_ "strings"

	"durn2.0/auth"
	db "durn2.0/database"
	"durn2.0/durn"
	"durn2.0/models"
	"durn2.0/util"
	_ "github.com/google/uuid"
)

// GetElections will fetch all elections in the system that the current
// user is authorized to view. Will include election info.
// Endpoint: GET /api/elections
func GetElections(res http.ResponseWriter, req *http.Request) {

}

// CreateElection creates a new election, given some initial data.
// Requires admin privileges
// Endpoint: POST /api/election/create
func CreateElection(res http.ResponseWriter, req *http.Request) {

}

// GetElectionInfo fetches all general info about a specific election,
// including: election name, candidates, when the election opens and closes,
// and whether the election is published.
// Endpoint: GET /api/election/{electionID}
func GetElectionInfo(res http.ResponseWriter, req *http.Request) {

}

// ModifyElection can change the following data for an election:
// Name, candidates, and open and close time.
// Requires admin privileges.
// Possibly: CAN'T CHANGE (CANDIDATES OF?) ONGOING ELECTION??
// Endpoint: PUT /api/election/{electionID}
func ModifyElection(res http.ResponseWriter, req *http.Request) {

}

// PublishElection marks an elections as published, marking it as ready
// for voting. Requires admin privileges.
// Endpoint: PUT /api/election/{electionID}/publish
func PublishElection(res http.ResponseWriter, req *http.Request) {

}

// UnpublishElection is the inverse of PublishElection
// Requires admin privileges.
// Endpoint: PUT /api/election/{electionID}/unpublish
func UnpublishElection(res http.ResponseWriter, req *http.Request) {

}

// CloseElection finalizes the result of the election, marking that
// it can be opened for public verification.
// Requires admin privileges.
// Endpoint: PUT /api/election/{electionID}/close
func CloseElection(res http.ResponseWriter, req *http.Request) {

}

// CastVote cast a vote in the specified election for the
// currently authenticated user.
// Endpoint: POST /api/elections/{electionID}/vote
func CastVote(res http.ResponseWriter, req *http.Request) {

}

// GetElectionVotes fetches a list of all votes in the selected election.
// Requires admin privileges while the election is open,
// but will be open to all once it has ended, to allow public verification
// Endpoint: GET /api/election/{electionID}/votes
func GetElectionVotes(res http.ResponseWriter, req *http.Request) {

}

// GetElectionVoteHashes fetches a list of all vote-hashes in the selected
// election. Requires admin privileges while the election is open,
// but will be open to all once it has ended, to allow public verification
// Endpoint: GET /api/election/{electionID}/votes/hashes
func GetElectionVoteHashes(res http.ResponseWriter, req *http.Request) {

}

// CountElectionVotes runs the counting algorithm for the specified
// election. Will return all intermediate steps of the algorithm to allow
// showing the process in frontend.
// Requires admin privileges while the election is open,
// but will be open to all once it has ended, to allow public verification
// Endpoint: GET /api/election/{electionID}/votes/count
func CountElectionVotes(res http.ResponseWriter, req *http.Request) {

}

// GetAllCandidates fetches info about all candidates that the user is
// authorized to view, i.e. in elections that the user can view.
// Endpoint: GET /api/candidates
func GetAllCandidates(res http.ResponseWriter, req *http.Request) {

}

// CreateCandidate creates a new candidate, given:
// Their name and a link to a candidate presentation.
// Requires admin privileges
// Endpoint: POST /api/candidates/create
func CreateCandidate(res http.ResponseWriter, req *http.Request) {

}

// GetCandidate fetches info about a candidate, if the user is
// authorized to view it.
// Endpoint: GET /api/candidate/{candidateID}
func GetCandidate(res http.ResponseWriter, req *http.Request) {

}

// ModifyCandidate changes the following info for a candidate:
// their name and the link to their candidate presentation.
// Requires admin privileges
// Endpoint: PUT /api/candidate/{candidateID}
func ModifyCandidate(res http.ResponseWriter, req *http.Request) {

}

// DeleteCandidate removes a candidate from the system
// Requires admin privileges.
// Should possibly only work if the user is in no elections?
// TBD if this functionality should be in the system
// Endpoint: DELETE /api/candidate/{candidateID}
func DeleteCandidate(res http.ResponseWriter, req *http.Request) {

}

// GetValidVoters fetches a list of all current users which are
// authorized to vote in elections.
// Requires admin privileges
// Endpoint: GET /api/voters
func GetValidVoters(res http.ResponseWriter, req *http.Request) {
	if !auth.IsAuthenticated(req.Context()) {
		err := util.AuthenticationError("Not authorized")
		util.RequestError(req.Context(), res, err)
		return
	}

	if err := auth.IsAuthorized(req.Context(), "viewAdmin"); err != nil {
		util.RequestError(req.Context(), res, err)
		return
	}

	voters, err := db.QueryAllVoters()
	if err != nil {
		util.RequestError(req.Context(), res, util.ServerError("Database query failed"))
		return
	}

	var response_data struct {
		Voters []models.Voter
	}

	response_data.Voters = voters
	err = util.WriteJson(res, response_data)
	if err != nil {
		util.RequestError(req.Context(), res, err)
		return
	}
}

// AddValidVoters adds a list of users(emails) to the list of valid voters
// Duplicates of users already in the system will be ignored
// Requires admin privileges
// Endpoint: PUT /api/voters/add
func AddValidVoters(res http.ResponseWriter, req *http.Request) {
	if !auth.IsAuthenticated(req.Context()) {
		err := util.AuthenticationError("Not authorized")
		util.RequestError(req.Context(), res, err)
		return
	}

	if err := auth.IsAuthorized(req.Context(), "viewAdmin"); err != nil {
		util.RequestError(req.Context(), res, err)
		return
	}

	var request_data struct {
		Voters []models.Voter
	}

	err := util.ReadJson(req, request_data)
	if err != nil {
		util.RequestError(req.Context(), res, err)
		return
	}

	err = durn.AddValidVoters(request_data.Voters)

	if err != nil {
		util.RequestError(req.Context(), res, err)
		return
	}
}

// RemoveValidVoters removes users from the list of valid voters.
// Users not in the list will be ignored
// Requires admin privileges
// Endpoint: PUT /api/voters/remove
func RemoveValidVoters(res http.ResponseWriter, req *http.Request) {

}

// GetLogs returns a list of all voting events that has occurred in the
// system, in order to make verifying a proper voting procedure easier
// Requires admin privileges
// Endpoint: GET /api/history
func GetLogs(res http.ResponseWriter, req *http.Request) {

}

// NukeSystem resets the system back to a clean state, removing all
// traces of all currently available elections
// Requires admin privileges
// Endpoint: PUT /api/reset-system
func NukeSystem(res http.ResponseWriter, req *http.Request) {

}
