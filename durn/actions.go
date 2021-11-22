package durn

import (
	"context"
	"fmt"
	"regexp"

	db "durn2.0/database"
	"durn2.0/models"
	rl "durn2.0/requestLog"
	"durn2.0/util"
)

// AddValidVoters checks that provided voters are entered as valid emails, and if that
// is the case inserts them into the database
func AddValidVoters(ctx context.Context, voters []models.Voter) error {
	mailregex := "[a-zA-Z]+@kth\\.se"
	for _, voter := range voters {
		matches, err := regexp.MatchString(mailregex, string(voter))
		if !matches {
			return util.BadRequestError(fmt.Sprintf("Trying to add email address '%s'", voter))
		}
		if err != nil {
			rl.Warning(ctx, err.Error())
			return util.ServerError("An internal server error occurred")
		}
	}

	err := db.InsertVoters(voters)
	if err != nil {
		rl.Warning(ctx, err.Error())
		return util.ServerError("Error while inserting into database")
	}

	return nil
}
