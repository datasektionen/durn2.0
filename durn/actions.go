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
// if that is the case inserts them into the database.
// Returns all voters for which it failed
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

	result := dbConn.Create(&dbVoters)
	if result.Error != nil {
		rl.Warning(ctx, result.Error.Error())
		return nil, util.ServerError("An internal server error occurred")
	}

	return failedVoters, nil
}

// GetAllValidVoters fetches all valid voters from the database
func GetAllValidVoters(ctx context.Context) ([]models.Voter, error) {
	dbConn := db.TakeDB()
	defer db.ReleaseDB()

	var voters []models.Voter
	var validVoters []db.ValidVoter
	result := dbConn.Find(&validVoters)

	if result.Error != nil {
		rl.Warning(ctx, result.Error.Error())
		return nil, util.ServerError("An internal server error occurred")
	}

	for _, voter := range validVoters {
		voters = append(voters, models.Voter(voter.Email))
	}

	return voters, nil
}
