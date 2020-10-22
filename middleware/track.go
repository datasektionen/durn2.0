package middleware

import (
	"context"
	"github.com/google/uuid"
	"net/http"
)

const reqIdKey string = "reqId"

func Track(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		ctx := req.Context()

		id := uuid.New()
		ctx = context.WithValue(ctx, reqIdKey, id)
		req = req.WithContext(ctx)

		next.ServeHTTP(res, req)
	})
}

func ReqId(r *http.Request) (uuid.UUID, bool) {
	reqId, ok := r.Context().Value(reqIdKey).(uuid.UUID)
	return reqId, ok
}
