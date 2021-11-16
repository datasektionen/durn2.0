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

//POST /api/election/{electionID}
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

// func GetElectionIds(res http.ResponseWriter, req *http.Request) {
// 	queriedStates := req.URL.Query().Get("states")
// 	if queriedStates == "" {
// 		queriedStates = "voting"
// 	}

// 	electionIds := make([]uuid.UUID, 0)

// 	if strings.Contains(queriedStates, "unpublished") {
// 		unpublishedElectionIds, err := durn.GetUnpublishedElectionIds(req.Context())
// 		if err != nil {
// 			util.RequestError(req.Context(), res, err)
// 		}

// 		electionIds = append(electionIds, unpublishedElectionIds...)
// 	}
// 	if strings.Contains(queriedStates, "voting") {
// 		unpublishedElectionIds, err := durn.GetVotingElectionIds(req.Context())
// 		if err != nil {
// 			util.RequestError(req.Context(), res, err)
// 		}

// 		electionIds = append(electionIds, unpublishedElectionIds...)
// 	}
// 	if strings.Contains(queriedStates, "closed") {
// 		unpublishedElectionIds, err := durn.GetClosedElectionIds(req.Context())
// 		if err != nil {
// 			util.RequestError(req.Context(), res, err)
// 		}

// 		electionIds = append(electionIds, unpublishedElectionIds...)
// 	}

// 	err := util.WriteJson(res, electionIds)
// 	if err != nil {
// 		util.RequestError(req.Context(), res, err)
// 	}
// }

// func CreateElection(res http.ResponseWriter, req *http.Request) {
// 	// var data struct {
// 	// 	Name         string
// 	// 	Alternatives []durn.Alternative
// 	// }

// 	// err := util.ReadJson(req, &data)
// 	// if err != nil {
// 	// 	util.RequestError(req.Context(), res, err)
// 	// }

// 	// err = durn.CreateElection(req.Context(), data.Name, data.Alternatives)
// 	// if err != nil {
// 	// 	util.RequestError(req.Context(), res, err)
// 	// 	return
// 	// }
// }

// func PublishElection(res http.ResponseWriter, req *http.Request) {
// 	electionId, err := util.GetPathUuid(req, "electionId")
// 	if err != nil {
// 		util.RequestError(req.Context(), res, err)
// 		return
// 	}

// 	err = durn.PublishElection(req.Context(), *electionId)
// 	if err != nil {
// 		util.RequestError(req.Context(), res, err)
// 		return
// 	}
// }

// func CloseElection(res http.ResponseWriter, req *http.Request) {
// 	electionId, err := util.GetPathUuid(req, "electionId")
// 	if err != nil {
// 		util.RequestError(req.Context(), res, err)
// 		return
// 	}

// 	err = durn.CloseElection(req.Context(), *electionId)
// 	if err != nil {
// 		util.RequestError(req.Context(), res, err)
// 		return
// 	}
// }

// func AddEligibleVoters(res http.ResponseWriter, req *http.Request) {
// 	electionId, err := util.GetPathUuid(req, "electionId")
// 	if err != nil {
// 		util.RequestError(req.Context(), res, err)
// 		return
// 	}

// 	var data struct {
// 		Voters []string
// 	}

// 	err = util.ReadJson(req, &data)
// 	if err != nil {
// 		util.RequestError(req.Context(), res, err)
// 		return
// 	}

// 	err = durn.AddEligibleVoters(req.Context(), *electionId, data.Voters)
// 	if err != nil {
// 		util.RequestError(req.Context(), res, err)
// 	}
// }

// func GetEligibleVoters(res http.ResponseWriter, req *http.Request) {
// 	electionId, err := util.GetPathUuid(req, "electionId")
// 	if err != nil {
// 		util.RequestError(req.Context(), res, err)
// 		return
// 	}

// 	voters, err := durn.GetEligibleVoters(req.Context(), *electionId)
// 	if err != nil {
// 		util.RequestError(req.Context(), res, err)
// 	}

// 	err = util.WriteJson(res, voters)
// 	if err != nil {
// 		util.RequestError(req.Context(), res, err)
// 	}
// }

// func CastVote(res http.ResponseWriter, req *http.Request) {
// 	// electionId, err := util.GetPathUuid(req, "electionId")
// 	// if err != nil {
// 	// 	util.RequestError(req.Context(), res, err)
// 	// 	return
// 	// }

// 	// var data struct {
// 	// 	Alternative durn.Alternative
// 	// }

// 	// err = util.ReadJson(req, &data)
// 	// if err != nil {
// 	// 	util.RequestError(req.Context(), res, err)
// 	// 	return
// 	// }

// 	// voteId, err := durn.CastVote(req.Context(), *electionId, data.Alternative)
// 	// if err != nil {
// 	// 	util.RequestError(req.Context(), res, err)
// 	// 	return
// 	// }

// 	// err = util.WriteJson(res, *voteId)
// 	// if err != nil {
// 	// 	util.RequestError(req.Context(), res, err)
// 	// }
// }

// func GetVotes(res http.ResponseWriter, req *http.Request) {
// 	electionId, err := util.GetPathUuid(req, "electionId")
// 	if err != nil {
// 		util.RequestError(req.Context(), res, err)
// 		return
// 	}

// 	votes, err := durn.GetVotes(req.Context(), *electionId)
// 	if err != nil {
// 		util.RequestError(req.Context(), res, err)
// 	}

// 	err = util.WriteJson(res, votes)
// 	if err != nil {
// 		util.RequestError(req.Context(), res, err)
// 	}
// }
