package middleware

import (
	"context"
	"github.com/google/uuid"
	"net/http"
)

// Key for request id in context
const reqIdKey string = "reqId"

// Middleware for attaching a request id to a request
func Track(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		ctx := req.Context()

		id := uuid.New()
		ctx = context.WithValue(ctx, reqIdKey, id)
		req = req.WithContext(ctx)

		next.ServeHTTP(res, req)
	})
}

// Get the request ID from a context
func ReqId(ctx context.Context) (uuid.UUID, bool) {
	reqId, ok := ctx.Value(reqIdKey).(uuid.UUID)
	return reqId, ok
}

// Get the request ID from a context
// Panic upon failure
func MustReqId(ctx context.Context) uuid.UUID {
	reqId, ok := ReqId(ctx)

	if !ok {
		panic("Request ID is missing from context")
	}

	return reqId
}
