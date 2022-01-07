package auth

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"durn2.0/conf"
	rl "durn2.0/requestLog"
	"durn2.0/util"
)

const apiUrlFormat string = "https://pls.datasektionen.se/api/user/%s/durn/%s"

// Permission names
const (
	AdminView   = "viewAdmin"
	AdminModify = "modifyAdmin"
)

func IsAuthorized(ctx context.Context, permission string) error {

	c := conf.ReadConfiguration()
	if c.SkipAuth {
		return nil
	}

	user, ok := util.User(ctx)
	if !ok {
		return errors.New("user not found in context")
	}

	url := fmt.Sprintf(apiUrlFormat, user, permission)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != 200 {
		return errors.New("non-ok (200) response from pls")
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var authorized bool
	err = json.Unmarshal(body, &authorized)
	if err != nil {
		return err
	}

	if !authorized {
		return util.AuthorizationError(fmt.Sprintf("%s not authorized for %s", user, permission))
	}

	rl.Info(ctx, fmt.Sprintf("%s verified to be authorized for %s", user, permission))

	return nil
}
