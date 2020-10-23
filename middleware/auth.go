package middleware

import (
	"context"
	"net/http"
)

// Key for user id in context
const userKey string = "user"

// Middleware for authenticating the client
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

// Get the user id from a context
func User(ctx context.Context) (string, bool) {
	user, ok := ctx.Value(userKey).(string)
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
