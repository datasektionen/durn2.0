package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"durn2.0/auth"
	_ "durn2.0/auth"
	"durn2.0/conf"
	db "durn2.0/database"
	"durn2.0/handler"
	mw "durn2.0/middleware"
	rl "durn2.0/requestLog"
	"durn2.0/util"

	"github.com/gorilla/mux"
)

func main() {
	c := conf.ReadConfiguration()

	rl.SetPrefixFn(func(ctx context.Context) string {
		reqId, ok := util.ReqId(ctx)
		if ok {
			return fmt.Sprintf("%s", reqId)
		} else {
			return fmt.Sprintf("missing")
		}
	})

	authenticator := auth.AuthenticationMiddleware{
		ApiKey: c.LoginApiKey,
	}

	db.CreateDBConnection()

	r := mux.NewRouter()
	r.Use(mw.Track)
	r.Use(mw.RequestLog)
	r.Use(mw.ResponseLog)

	a := r.PathPrefix("/api").Subrouter()
	a.Use(authenticator.Middleware)

	s := a.PathPrefix("/elections").Subrouter()
	s.Methods("GET").HandlerFunc(handler.GetElections)

	s = a.PathPrefix("/election").Subrouter()
	s.Path("/create").Methods("POST").HandlerFunc(handler.CreateElection)

	s = a.PathPrefix("/election/{electionId}").Subrouter()
	s.Methods("GET").HandlerFunc(handler.GetElectionInfo)
	s.Methods("PUT").HandlerFunc(handler.ModifyElection)
	s.Path("/publish").Methods("PUT").HandlerFunc(handler.PublishElection)
	s.Path("/unpublish").Methods("PUT").HandlerFunc(handler.UnpublishElection)
	s.Path("/close").Methods("PUT").HandlerFunc(handler.CloseElection)
	s.Path("/vote").Methods("POST").HandlerFunc(handler.CastVote)
	s.Path("/votes").Methods("GET").HandlerFunc(handler.GetElectionVotes)
	s.Path("/votes/hashes").Methods("GET").HandlerFunc(handler.GetElectionVoteHashes)
	s.Path("/votes/count").Methods("GET").HandlerFunc(handler.CountElectionVotes)

	s = a.PathPrefix("/candidates").Subrouter()
	s.Methods("GET").HandlerFunc(handler.GetAllCandidates)

	s = a.PathPrefix("/candidate").Subrouter()
	s.Path("/create").Methods("POST").HandlerFunc(handler.CreateCandidate)

	s = a.PathPrefix("/candidate/{candidateID}").Subrouter()
	s.Methods("GET").HandlerFunc(handler.GetCandidate)
	s.Methods("PUT").HandlerFunc(handler.ModifyCandidate)
	// s.Methods("DELETE").HandlerFunc(handler.DeleteCandidate)

	s = a.PathPrefix("/voters").Subrouter()
	s.Methods("GET").HandlerFunc(handler.GetValidVoters)
	s.Path("/add").Methods("PUT").HandlerFunc(handler.AddValidVoters)
	s.Path("/remove").Methods("PUT").HandlerFunc(handler.RemoveValidVoters)

	s = a.PathPrefix("/history").Subrouter()
	s.Methods("GET").HandlerFunc(handler.GetLogs)

	s = a.PathPrefix("/reset-system").Subrouter()
	s.Methods("PUT").HandlerFunc(handler.NukeSystem)

	server := http.Server{
		Addr:    c.Addr,
		Handler: r,
	}

	log.Printf("Starting server on %s\n", c.Addr)
	log.Fatal(server.ListenAndServe())
}
