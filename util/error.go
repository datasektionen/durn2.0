package util

import (
	"context"
	rl "durn2.0/requestLog"
	"net/http"
)

func RequestError(ctx context.Context, res http.ResponseWriter, err error) {
	var status int

	if err, ok := err.(ApplicationError); ok {
		status = err.StatusCode()

		for key, val := range err.Headers() {
			res.Header().Set(key, val)
		}
	} else {
		status = http.StatusInternalServerError
	}


	rl.Warning(ctx, err.Error())
	res.WriteHeader(status)
}

type ApplicationError interface {
	error
	StatusCode() int
	Headers() map[string]string
}

type AuthenticationError string

func (a AuthenticationError) Error() string {
	return string(a)
}

func (a AuthenticationError) StatusCode() int {
	return http.StatusUnauthorized
}

func (a AuthenticationError) Headers() map[string]string {
	return map[string]string{
		"WWW-Authenticate": "Bearer, " +
			"error=\"invalid_token\", " +
			"error_description=\"Invalid or expired access token\"",
	}
}

type AuthorizationError string

func (a AuthorizationError) Error() string {
	return string(a)
}

func (a AuthorizationError) StatusCode() int {
	panic("implement me")
}

func (a AuthorizationError) Headers() map[string]string {
	panic("implement me")
}

type BadRequestError string

func (b BadRequestError) Error() string {
	panic("implement me")
}

func (b BadRequestError) StatusCode() int {
	panic("implement me")
}

func (b BadRequestError) Headers() map[string]string {
	panic("implement me")
}


