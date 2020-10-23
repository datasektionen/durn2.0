package middleware

import (
	rl "durn2.0/requestLog"
	"fmt"
	"github.com/felixge/httpsnoop"
	"net/http"
)

// Middleware for logging info about request.
//
// Will log upon a request is received.
// Will log the request method and path.
func RequestLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		method := req.Method
		path := req.URL.Path

		rl.Info(req, fmt.Sprintf("%s %s", method, path))

		next.ServeHTTP(res, req)
	})
}

// Middleware for logging info about response.
//
// Will log upon once response is ready.
// Will log:
// * the status code of the response,
// * how long the request took to process in milliseconds, and
// * how many bytes were written.
func ResponseLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		m := httpsnoop.CaptureMetrics(next, res, req)

		txt := fmt.Sprintf(
			"status: %d, duration: %dms, body: %d bytes",
			m.Code,
			m.Duration.Milliseconds(),
			m.Written,
		)

		rl.Info(req, txt)
	})
}