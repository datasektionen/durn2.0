package middleware

import (
	"context"
	"durn2.0/util"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type AuthenticationMiddleware struct {
	ApiKey string
}

func (a *AuthenticationMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		ctx := req.Context()

		cookie, err := req.Cookie("sessionID")
		if err != nil {
			util.RequestError(
				res, req, http.StatusUnauthorized, err,
				"Login token cookie missing from request",
			)
			return
		}
		
		token := cookie.Value
		apiUrl := fmt.Sprintf("https://login.datasektionen.se/verify/%s.json?api_key=%s", token, a.ApiKey)
		
		resp, err := http.Get(apiUrl)
		if err != nil {
			util.RequestError(
				res, req, http.StatusInternalServerError, err,
				"Error sending request to verify login token",
			)
			return
		}

		// Check that response is ok
		if resp.StatusCode - resp.StatusCode % 200 != 200 {
			util.RequestError(
				res, req, http.StatusInternalServerError, nil,
				"Non-OK status code while verifying token",
			)
			return
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			util.RequestError(
				res, req, http.StatusInternalServerError, err,
				"Error reading response to token verification",
			)
			return
		}

		var data loginVerificationData
		err = json.Unmarshal(body, &data)
		if err != nil {
			util.RequestError(
				res, req, http.StatusInternalServerError, err,
				"Error unmarshalling json in token verification",
			)
		}
		
		ctx = context.WithValue(ctx, util.UserKey, data.KthID)
		req = req.WithContext(ctx)

		next.ServeHTTP(res, req)
	})
}

type loginVerificationData struct {
	Email string `json:"email"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	KthID string `json:"ugkthid"`
	UserName string `json:"user"`
}