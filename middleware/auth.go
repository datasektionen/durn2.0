package middleware

import (
	"context"
	"net/http"
)

const userKey string = "user"

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		ctx := req.Context()

		// todo: acquire user properly with login2
		user := "jespel"

		ctx = context.WithValue(ctx, userKey, user)
		req = req.WithContext(ctx)

		next.ServeHTTP(res, req)
	})
}

func User(r *http.Request) (string, bool) {
	user, ok := r.Context().Value(userKey).(string)
	return user, ok
}
