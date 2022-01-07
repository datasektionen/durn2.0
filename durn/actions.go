package durn

import (
	"context"
	"regexp"

	db "durn2.0/database"
	"durn2.0/models"
	rl "durn2.0/requestLog"
	"durn2.0/util"
)

// AddValidVoters checks that provided voters are entered as valid emails, and
// if that is the case inserts them into the database
func AddValidVoters(ctx context.Context, voters []models.Voter) ([]models.Voter, error) {
	mailRegex := "[^@]+@kth\\.se"
	dbVoters := []db.ValidVoter{}
	failedVoters := []models.Voter{}

	for _, voter := range voters {
		matches, err := regexp.MatchString(mailRegex, string(voter))
		if err != nil {
			rl.Warning(ctx, err.Error())
			return nil, util.ServerError("An internal server error occurred")
		}

		if !matches {
			failedVoters = append(failedVoters, voter)
		} else {
			dbVoters = append(dbVoters, db.ValidVoter{
				Email: string(voter),
			})
		}
	}

	dbConn := db.TakeDB()
	defer db.ReleaseDB()

	dbConn.Create(&dbVoters)

	return failedVoters, nil
}
