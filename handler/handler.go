package handler

import (
	"fmt"
	"net/http"
	_ "strings"

	_ "durn2.0/durn"
	rl "durn2.0/requestLog"
	"durn2.0/util"
	_ "github.com/google/uuid"
)

func Login(res http.ResponseWriter, req *http.Request) {
	apiUrl := "https://login.datasektionen.se/login?callback="
	callbackUrl := fmt.Sprintf("http://%s/login-complete?token=", req.Host)
	redirectUrl := fmt.Sprintf("%s%s", apiUrl, callbackUrl)

	res.Header().Set("Location", redirectUrl)
	res.WriteHeader(http.StatusTemporaryRedirect)
}

func LoginComplete(res http.ResponseWriter, req *http.Request) {
	rl.Info(req.Context(), "Login complete")

	token, ok := req.URL.Query()["token"]
	if !ok || len(token) == 0 {
		util.RequestError(req.Context(), res, util.BadRequestError("missing token parameter from request"))
	}

	_, _ = res.Write([]byte(token[0]))
}

// GET /api/elections
func GetElections(res http.ResponseWriter, req *http.Request) {

}

// POST /api/election/create
func CreateElection(res http.ResponseWriter, req *http.Request) {

}

// GET /api/election/{electionID}
func GetElectionInfo(res http.ResponseWriter, req *http.Request) {

}

// POST /api/election/{electionID}
// CAN'T CHANGE (CANDIDATES OF?) ONGOING ELECTION??
func ModifyElection(res http.ResponseWriter, req *http.Request) {

}

// PUT /api/election/{electionID}/publish
func PublishElection(res http.ResponseWriter, req *http.Request) {

}

// PUT /api/election/{electionID}/close
func CloseElection(res http.ResponseWriter, req *http.Request) {

}

// POST /api/elections/{electionID}/vote
func CastVote(res http.ResponseWriter, req *http.Request) {

}

// GET /api/election/{electionID}/votes
// AVAILABLE FOR ALL AFTER ELECTION HAS CLOSED
func GetElectionVotes(res http.ResponseWriter, req *http.Request) {

}

// GET /api/election/{electionID}/votes/count
func CountElectionVotes(res http.ResponseWriter, req *http.Request) {

}

// GET /api/candidates
func GetAllCandidates(res http.ResponseWriter, req *http.Request) {

}

// POST /api/candidates/create
func CreateCandidate(res http.ResponseWriter, req *http.Request) {

}

// GET /api/candidate/{candidateID}
func GetCandidate(res http.ResponseWriter, req *http.Request) {

}

// POST /api/candidate/{candidateID}
func ModifyCandidate(res http.ResponseWriter, req *http.Request) {

}

// DELETE /api/candidate/{candidateID}
// ONLY IF NOT IN ANY ELECTIONS?
func DeleteCandidate(res http.ResponseWriter, req *http.Request) {

}

// GET /api/voters
func GetValidVoters(res http.ResponseWriter, req *http.Request) {

}

// PUT /api/voters/add
func AddValidVoters(res http.ResponseWriter, req *http.Request) {

}

// PUT /api/voters/remove
func RemoveValidVoters(res http.ResponseWriter, req *http.Request) {

}

// GET /api/history
func GetLogs(res http.ResponseWriter, req *http.Request) {

}

// PUT /api/clearDB
func NukeSystem(res http.ResponseWriter, req *http.Request) {

}
