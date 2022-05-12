package middleware

import (
	"context"
	"durn2.0/util"
	"github.com/google/uuid"
	"net/http"
)


// Middleware for attaching a request id to a request
func Track(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		ctx := req.Context()

		id := uuid.New()
		ctx = context.WithValue(ctx, util.ReqIdKey, id)
		req = req.WithContext(ctx)

		next.ServeHTTP(res, req)
	})
}


