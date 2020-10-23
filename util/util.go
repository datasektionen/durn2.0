package util

import (
	"context"
	rl "durn2.0/requestLog"
	"fmt"
	"github.com/google/uuid"
	"net/http"
)

// Key for user id in context
const UserKey string = "user"

// Key for request id in context
const ReqIdKey string = "reqId"

func RequestError(res http.ResponseWriter, req *http.Request, status int, err error, format string, v ...interface{}) {
	if err != nil {
		format = fmt.Sprintf("%s: %v", format, err)
	}

	desc := fmt.Sprintf(format, v...)

	rl.Warning(req, desc)
	res.WriteHeader(status)
}

// Get the request ID from a context
func ReqId(ctx context.Context) (uuid.UUID, bool) {
	reqId, ok := ctx.Value(ReqIdKey).(uuid.UUID)
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

// Get the user id from a context
func User(ctx context.Context) (string, bool) {
	user, ok := ctx.Value(UserKey).(string)
	return user, ok
}

// Get the user id from context
// Panics upon failure
func MustUser(ctx context.Context) string {
	user, ok := User(ctx)

	if !ok {
		panic("User ID is missing from context")
	}

	return user
}