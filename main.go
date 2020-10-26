package main

import (
	"context"
	"durn2.0/auth"
	"durn2.0/conf"
	"durn2.0/handler"
	mw "durn2.0/middleware"
	rl "durn2.0/requestLog"
	"durn2.0/util"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
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

	r := mux.NewRouter()
	r.Use(mw.Track)
	r.Use(mw.RequestLog)
	r.Use(mw.ResponseLog)

	o := r.PathPrefix("/").Subrouter()
	o.Path("/login").Methods("GET").HandlerFunc(handler.Login)
	o.Path("/login-complete").Methods("GET").HandlerFunc(handler.LoginComplete)

	a := r.PathPrefix("/api").Subrouter()
	a.Use(authenticator.Middleware)

	s := a.PathPrefix("/elections").Subrouter()
	s.Methods("GET").HandlerFunc(handler.GetElectionIds)
	s.Methods("POST").HandlerFunc(handler.CreateElection)

	s = a.PathPrefix("/election/{electionId}").Subrouter()
	s.Path("/vote").Methods("POST").HandlerFunc(handler.CastVote)
	s.Path("/voters").Methods("GET").HandlerFunc(handler.GetEligibleVoters)
	s.Path("/voters").Methods("PUT").HandlerFunc(handler.AddEligibleVoters)

	server := http.Server{
		Addr: c.Addr,
		Handler: r,
	}

	log.Printf("Starting server on %s\n", c.Addr)
	log.Fatal(server.ListenAndServe())
}
