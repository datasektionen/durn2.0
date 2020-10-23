package handler

import (
	"durn2.0/durn"
	rl "durn2.0/requestLog"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

func requestError(res http.ResponseWriter, req *http.Request, status int, err error, format string, v ...interface{}) {
	if err != nil {
		format = fmt.Sprintf("%s: %v", format, err)
	}

	desc := fmt.Sprintf(format, v...)

	rl.Warning(req, desc)
	res.WriteHeader(status)
}

func NotFound(res http.ResponseWriter, req *http.Request) {
	requestError(
		res, req, http.StatusNotFound, nil,
		"Not found",
	)
}

func Login(res http.ResponseWriter, req *http.Request) {
	apiUrl := "https://login.datasektionen.se/login?callback="
	callbackUrl := fmt.Sprintf("http://%s/login-complete?token=", req.Host)
	redirectUrl := fmt.Sprintf("%s%s", apiUrl, callbackUrl)

	res.Header().Set("Location", redirectUrl)
	res.WriteHeader(http.StatusTemporaryRedirect)
}

func LoginComplete(res http.ResponseWriter, req *http.Request) {
	rl.Info(req, "Login complete")

	token, ok := req.URL.Query()["token"]
	if !ok || len(token) == 0 {
		requestError(
			res, req, http.StatusBadRequest, nil,
			"Missing token parameter from request",
		)
	}

	cookie := http.Cookie{
		Name:       "sessionID",
		Value:      token[0],
		Secure:     false,
		HttpOnly:   true,
		SameSite:   http.SameSiteLaxMode,
	}

	http.SetCookie(res, &cookie)
}

func GetElections(res http.ResponseWriter, req *http.Request) {
	elections := durn.GetElections()

	data, err := json.Marshal(elections)
	if err != nil {
		requestError(
			res, req, http.StatusInternalServerError, err,
			"Error while marshal election data as JSON",
		)
		return
	}

	res.Header().Set("Content-Type", "application/json")

	_, _ = res.Write(data)
}

func CreateElection(res http.ResponseWriter, req *http.Request) {
	type createElectionData struct {
		Name         string
		Alternatives []durn.Alternative
	}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		requestError(
			res, req, http.StatusInternalServerError, err,
			"Error while reading request body",
		)
		return
	}

	var data createElectionData
	err = json.Unmarshal(body, &data)
	if err != nil {
		requestError(
			res, req, http.StatusBadRequest, err,
			"Request body could not be unmarshalled as JSON",
		)
		return
	}

	durn.CreateElection(req.Context(), data.Name, data.Alternatives)
}

func AddEligibleVoters(res http.ResponseWriter, req *http.Request) {
	type addEligibleVotersData struct {
		Voters []string
	}

	electionIdString, ok := mux.Vars(req)["electionId"]
	if !ok {
		requestError(
			res, req, http.StatusBadRequest, nil,
			"Request is missing election ID from path",
		)
		return
	}

	electionId, err := uuid.Parse(electionIdString)
	if err != nil {
		requestError(
			res, req, http.StatusBadRequest, err,
			"Given election ID cannot be parsed as UUID",
		)
		return
	}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		requestError(
			res, req, http.StatusInternalServerError, err,
			"Error while reading request body",
		)
		return
	}

	var data addEligibleVotersData
	err = json.Unmarshal(body, &data)
	if err != nil {
		requestError(
			res, req, http.StatusBadRequest, err,
			"Request body could not be unmarshalled as JSON",
		)
		return
	}

	durn.AddEligibleVoters(req.Context(), electionId, data.Voters)
}

func GetEligibleVoters(res http.ResponseWriter, req *http.Request) {
	electionIdString, ok := mux.Vars(req)["electionId"]
	if !ok {
		requestError(
			res, req, http.StatusBadRequest, nil,
			"Request is missing election ID from path",
		)
		return
	}

	electionId, err := uuid.Parse(electionIdString)
	if err != nil {
		requestError(
			res, req, http.StatusBadRequest, err,
			"Given election ID cannot be parsed as UUID",
		)
		return
	}

	voters, err := durn.GetEligibleVoters(req.Context(), electionId)

	data, err := json.Marshal(voters)
	if err != nil {
		requestError(
			res, req, http.StatusInternalServerError, err,
			"Error while marshal voter data as JSON",
		)
		return
	}

	res.Header().Set("Content-Type", "application/json")

	_, _ = res.Write(data)
}

func CastVote(res http.ResponseWriter, req *http.Request) {
	type castVoteData struct {
		Alternative durn.Alternative
	}

	electionIdString, ok := mux.Vars(req)["electionId"]
	if !ok {
		requestError(
			res, req, http.StatusBadRequest, nil,
			"Request is missing election ID from path",
		)
		return
	}

	electionId, err := uuid.Parse(electionIdString)
	if err != nil {
		requestError(
			res, req, http.StatusBadRequest, err,
			"Given election ID cannot be parsed as UUID",
		)
		return
	}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		requestError(
			res, req, http.StatusInternalServerError, err,
			"Error while reading request body",
		)
		return
	}

	var data castVoteData
	err = json.Unmarshal(body, &data)
	if err != nil {
		requestError(
			res, req, http.StatusBadRequest, err,
			"Request body could not be unmarshalled as JSON",
		)
		return
	}

	err = durn.CastVote(req.Context(), electionId, data.Alternative)
	if err != nil {
		requestError(
			res, req, http.StatusBadRequest, err,
			"Error casting vote",
		)
		return
	}
}
