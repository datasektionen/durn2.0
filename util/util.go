package util

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

// Key for user id in context
const UserKey string = "user"

// Key for request id in context
const ReqIdKey string = "reqId"

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

func GetPathUuid(req *http.Request, key string) (*uuid.UUID, error) {
	raw, ok := mux.Vars(req)[key]
	if !ok {
		return nil, BadRequestError(fmt.Sprintf("missing %s from path", key))
	}

	id, err := uuid.Parse(raw)
	if err != nil {
		return nil, err
	}

	return &id, nil
}

func ReadJson(req *http.Request, data interface{}) error {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, data)
	if err != nil {
		return err
	}

	return nil
}

func WriteJson(res http.ResponseWriter, data interface{}) error {
	marshalledData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	res.Header().Set("Content-Type", "application/json")
	_, err = res.Write(marshalledData)
	return err
}