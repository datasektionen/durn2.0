package handler

import (
	"durn2.0/durn"
	rl "durn2.0/requestLog"
	"durn2.0/util"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"strings"
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

func GetElectionIds(res http.ResponseWriter, req *http.Request) {
	queriedStates := req.URL.Query().Get("states")
	if queriedStates == "" {
		queriedStates = "voting"
	}

	electionIds := make([]uuid.UUID, 0)

	if strings.Contains(queriedStates, "unpublished") {
		unpublishedElectionIds, err := durn.GetUnpublishedElectionIds(req.Context())
		if err != nil {
			util.RequestError(req.Context(), res, err)
		}

		electionIds = append(electionIds, unpublishedElectionIds...)
	}
	if strings.Contains(queriedStates, "voting") {
		unpublishedElectionIds, err := durn.GetVotingElectionIds(req.Context())
		if err != nil {
			util.RequestError(req.Context(), res, err)
		}

		electionIds = append(electionIds, unpublishedElectionIds...)
	}
	if strings.Contains(queriedStates, "closed") {
		unpublishedElectionIds, err := durn.GetClosedElectionIds(req.Context())
		if err != nil {
			util.RequestError(req.Context(), res, err)
		}

		electionIds = append(electionIds, unpublishedElectionIds...)
	}

	err := util.WriteJson(res, electionIds)
	if err != nil {
		util.RequestError(req.Context(), res, err)
	}
}

func CreateElection(res http.ResponseWriter, req *http.Request) {
	var data struct {
		Name         string
		Alternatives []durn.Alternative
	}

	err := util.ReadJson(req, &data)
	if err != nil {
		util.RequestError(req.Context(), res, err)
	}

	err = durn.CreateElection(req.Context(), data.Name, data.Alternatives)
	if err != nil {
		util.RequestError(req.Context(), res, err)
		return
	}
}

func PublishElection(res http.ResponseWriter, req *http.Request) {
	electionId, err := util.GetPathUuid(req, "electionId")
	if err != nil {
		util.RequestError(req.Context(), res, err)
		return
	}

	err = durn.PublishElection(req.Context(), *electionId)
	if err != nil {
		util.RequestError(req.Context(), res, err)
		return
	}
}

func CloseElection(res http.ResponseWriter, req *http.Request) {
	electionId, err := util.GetPathUuid(req, "electionId")
	if err != nil {
		util.RequestError(req.Context(), res, err)
		return
	}

	err = durn.CloseElection(req.Context(), *electionId)
	if err != nil {
		util.RequestError(req.Context(), res, err)
		return
	}
}

func AddEligibleVoters(res http.ResponseWriter, req *http.Request) {
	electionId, err := util.GetPathUuid(req, "electionId")
	if err != nil {
		util.RequestError(req.Context(), res, err)
		return
	}

	var data struct {
		Voters []string
	}

	err = util.ReadJson(req, &data)
	if err != nil {
		util.RequestError(req.Context(), res, err)
		return
	}

	err = durn.AddEligibleVoters(req.Context(), *electionId, data.Voters)
	if err != nil {
		util.RequestError(req.Context(), res, err)
	}
}

func GetEligibleVoters(res http.ResponseWriter, req *http.Request) {
	electionId, err := util.GetPathUuid(req, "electionId")
	if err != nil {
		util.RequestError(req.Context(), res, err)
		return
	}

	voters, err := durn.GetEligibleVoters(req.Context(), *electionId)
	if err != nil {
		util.RequestError(req.Context(), res, err)
	}

	err = util.WriteJson(res, voters)
	if err != nil {
		util.RequestError(req.Context(), res, err)
	}
}

func CastVote(res http.ResponseWriter, req *http.Request) {
	electionId, err := util.GetPathUuid(req, "electionId")
	if err != nil {
		util.RequestError(req.Context(), res, err)
		return
	}

	var data struct {
		Alternative durn.Alternative
	}

	err = util.ReadJson(req, &data)
	if err != nil {
		util.RequestError(req.Context(), res, err)
		return
	}

	voteId, err := durn.CastVote(req.Context(), *electionId, data.Alternative)
	if err != nil {
		util.RequestError(req.Context(), res, err)
		return
	}

	err = util.WriteJson(res, *voteId)
	if err != nil {
		util.RequestError(req.Context(), res, err)
	}
}

func GetVotes(res http.ResponseWriter, req *http.Request) {
	electionId, err := util.GetPathUuid(req, "electionId")
	if err != nil {
		util.RequestError(req.Context(), res, err)
		return
	}

	votes, err := durn.GetVotes(req.Context(), *electionId)
	if err != nil {
		util.RequestError(req.Context(), res, err)
	}

	err = util.WriteJson(res, votes)
	if err != nil {
		util.RequestError(req.Context(), res, err)
	}
}
