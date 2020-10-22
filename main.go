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
		reqId, ok := mw.ReqId(req)
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

	s = r.PathPrefix("/election/{electionId}/vote").Subrouter()
	s.Methods("POST").HandlerFunc(castVote)

	server := http.Server{
		Addr: ":8080",
		Handler: r,
	}

	log.Fatal(server.ListenAndServe())
}

func getElections(res http.ResponseWriter, _ *http.Request) {
	elections := durn.GetElections()
	data, _ := json.Marshal(elections)
	_, _ = res.Write(data)
}

func createElection(_ http.ResponseWriter, req *http.Request) {
	type createElectionData struct {
		Name string
		Alternatives []durn.Alternative
	}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	var data createElectionData
	err = json.Unmarshal(body, &data)
	if err != nil {
		panic(err)
	}

	durn.CreateElection(data.Name, data.Alternatives)
}

func castVote(_ http.ResponseWriter, req *http.Request) {
	type castVoteData struct {
		Alternative durn.Alternative
	}

	electionIdString, ok := mux.Vars(req)["electionId"]
	if !ok {
		panic("missing election id")
	}

	electionId, err := uuid.Parse(electionIdString)
	if err != nil {
		panic(err)
	}

	user, ok := mw.User(req)
	if !ok {
		panic("missing user from req")
	}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	var data castVoteData
	err = json.Unmarshal(body, &data)
	if err != nil {
		panic(err)
	}

	err = durn.CastVote(electionId, user, data.Alternative)
	if err != nil {
		panic(err)
	}
}
