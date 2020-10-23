package main

import (
	"durn2.0/durn"
	mw "durn2.0/middleware"
	rl "durn2.0/requestLog"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	rl.SetPrefixFn(func(req *http.Request) string {
		reqId, ok := mw.ReqId(req.Context())
		if ok {
			return fmt.Sprintf("%s", reqId)
		} else {
			return fmt.Sprintf("missing")
		}
	})

	r.Use(mw.Track)
	r.Use(mw.Authenticate)
	r.Use(mw.RequestLog)
	r.Use(mw.ResponseLog)

	s := r.PathPrefix("/elections").Subrouter()
	s.Methods("GET").HandlerFunc(getElections)
	s.Methods("POST").HandlerFunc(createElection)

	s = r.PathPrefix("/election/{electionId}").Subrouter()
	s.Path("/vote").Methods("POST").HandlerFunc(castVote)
	s.Path("/voters").Methods("GET").HandlerFunc(getEligibleVoters)
	s.Path("/voters").Methods("PUT").HandlerFunc(addEligibleVoters)

	server := http.Server{
		Addr: ":8080",
		Handler: r,
	}

	log.Fatal(server.ListenAndServe())
}

func getElections(res http.ResponseWriter, req *http.Request) {
	elections := durn.GetElections()

	data, err := json.Marshal(elections)
	if err != nil {
		rl.Warning(req, fmt.Sprintf("Error while marshal election data as JSON: %v", err))
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")

	_, _ = res.Write(data)
}

func createElection(res http.ResponseWriter, req *http.Request) {
	type createElectionData struct {
		Name string
		Alternatives []durn.Alternative
	}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		rl.Warning(req, "Error while reading request body")
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	var data createElectionData
	err = json.Unmarshal(body, &data)
	if err != nil {
		rl.Warning(req, "Request body could not be unmarshalled as JSON")
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	durn.CreateElection(req.Context(), data.Name, data.Alternatives)
}

func addEligibleVoters(res http.ResponseWriter, req *http.Request) {
	type addEligibleVotersData struct {
		Voters []string
	}

	electionIdString, ok := mux.Vars(req)["electionId"]
	if !ok {
		rl.Warning(req, "Request is missing election ID from path")
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	electionId, err := uuid.Parse(electionIdString)
	if err != nil {
		rl.Warning(req, "Given election ID cannot be parsed as UUID")
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		rl.Warning(req, "Error while reading request body")
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	var data addEligibleVotersData
	err = json.Unmarshal(body, &data)
	if err != nil {
		rl.Warning(req, "Request body could not be unmarshalled as JSON")
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	durn.AddEligibleVoters(req.Context(), electionId, data.Voters)
}

func getEligibleVoters(res http.ResponseWriter, req *http.Request) {
	electionIdString, ok := mux.Vars(req)["electionId"]
	if !ok {
		rl.Warning(req, "Request is missing election ID from path")
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	electionId, err := uuid.Parse(electionIdString)
	if err != nil {
		rl.Warning(req, "Given election ID cannot be parsed as UUID")
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	voters, err := durn.GetEligibleVoters(req.Context(), electionId)

	data, err := json.Marshal(voters)
	if err != nil {
		rl.Warning(req, fmt.Sprintf("Error while marshal voter data as JSON: %v", err))
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")

	_, _ = res.Write(data)
}

func castVote(res http.ResponseWriter, req *http.Request) {
	type castVoteData struct {
		Alternative durn.Alternative
	}

	electionIdString, ok := mux.Vars(req)["electionId"]
	if !ok {
		rl.Warning(req, "Request is missing election ID from path")
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	electionId, err := uuid.Parse(electionIdString)
	if err != nil {
		rl.Warning(req, "Given election ID cannot be parsed as UUID")
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		rl.Warning(req, "Error while reading request body")
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	var data castVoteData
	err = json.Unmarshal(body, &data)
	if err != nil {
		rl.Warning(req, "Request body could not be unmarshalled as JSON")
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	err = durn.CastVote(req.Context(), electionId, data.Alternative)
	if err != nil {
		rl.Warning(req, fmt.Sprintf("Error casting vote: %v\n", err))
		res.WriteHeader(http.StatusBadRequest)
		return
	}
}
