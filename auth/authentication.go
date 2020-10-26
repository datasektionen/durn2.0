package auth

import (
	"context"
	rl "durn2.0/requestLog"
	"durn2.0/util"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const loginApiUrlFormat string = "https://login.datasektionen.se/verify/%s.json?api_key=%s"

type AuthenticationError string

func (a AuthenticationError) Error() string {
	return string(a)
}

type AuthenticationMiddleware struct {
	ApiKey string
}

func (a *AuthenticationMiddleware) authenticate(ctx context.Context, token string) (*AuthenticatedUser, error) {
	url := fmt.Sprintf(loginApiUrlFormat, token, a.ApiKey)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		if res.StatusCode == http.StatusNotFound {
			return nil, AuthenticationError("could not verify token with login")
		}
		return nil, errors.New("non-ok (200) response from login")
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var data AuthenticatedUser
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (a *AuthenticationMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		ctx := req.Context()

		authHeader := req.Header.Get("Authorization")
		token := strings.Replace(authHeader, "Bearer ", "", 1)

		user, err := a.authenticate(ctx, token)
		if err != nil {
			if authErr, ok := err.(AuthenticationError); ok {
				res.Header().Set("WWW-Authenticate",
					"Bearer," +
					"error=\"invalid_token\"," +
					"error_description=\"Invalid or expired access token\"")
				util.RequestError(res, req, http.StatusUnauthorized, authErr,
					"User not authenticated",
				)
			} else {
				util.RequestError(
					res, req, http.StatusInternalServerError, err,
					"Error while verifying user",
				)
			}
			return
		}

		rl.Info(req.Context(), fmt.Sprintf("Authenticated client with login"))

		ctx = context.WithValue(ctx, util.UserKey, user.UserName)
		req = req.WithContext(ctx)

		next.ServeHTTP(res, req)
	})
}

type AuthenticatedUser struct {
	Email string `json:"email"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	KthID string `json:"ugkthid"`
	UserName string `json:"user"`
}
