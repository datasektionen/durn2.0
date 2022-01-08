package durn

import (
	"context"
	"regexp"

	"gorm.io/gorm/clause"

	db "durn2.0/database"
	"durn2.0/models"
	rl "durn2.0/requestLog"
	"durn2.0/util"
)

// AddValidVoters checks that provided voters are entered as valid emails, and
// if that is the case inserts them into the database.
// Returns all voters for which it failed
func AddValidVoters(ctx context.Context, voters []models.Voter) ([]models.Voter, error) {
	dbConn := db.TakeDB()
	defer db.ReleaseDB()

	mailRegex := "[^@]+@kth\\.se"
	var votersToInsert []db.ValidVoter
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
			votersToInsert = append(votersToInsert, db.ValidVoter{
				Email: string(voter),
			})
		}
	}

	if len(votersToInsert) > 0 {
		if result := dbConn.Clauses(
			clause.OnConflict{DoNothing: true},
		).Create(&votersToInsert); result.Error != nil {
			rl.Warning(ctx, result.Error.Error())
			return nil, util.ServerError("An internal server error occurred")
		}
	}

	return failedVoters, nil
}

// GetAllValidVoters fetches all valid voters from the database
func GetAllValidVoters(ctx context.Context) ([]models.Voter, error) {
	dbConn := db.TakeDB()
	defer db.ReleaseDB()

	var validVoters []db.ValidVoter
	voters := []models.Voter{}

	if result := dbConn.Find(&validVoters); result.Error != nil {
		rl.Warning(ctx, result.Error.Error())
		return nil, util.ServerError("An internal server error occurred")
	}

	for _, voter := range validVoters {
		voters = append(voters, models.Voter(voter.Email))
	}

	return voters, nil
}

func DeleteValidVoters(ctx context.Context, voters []models.Voter) error {
	dbConn := db.TakeDB()
	defer db.ReleaseDB()
	var validVoters []db.ValidVoter

	if result := dbConn.Delete(&validVoters, voters); result.Error != nil {
		rl.Warning(ctx, result.Error.Error())
		return util.ServerError("An internal server error occurred")
	}

	return nil
}
