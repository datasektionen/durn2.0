package database

import (
	"fmt"
	"sync"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"durn2.0/conf"
	"durn2.0/models"
)

var mutex sync.Mutex
var db *gorm.DB

// CreateDBConnection Initializes the database connection using
// the connection details specified in env-vars
func CreateDBConnection() error {
	c := conf.ReadConfiguration()
	mutex.Lock()
	defer mutex.Unlock()

	dsn := fmt.Sprintf("host=%s port=%d user='%s' password='%s' dbname='%s' sslmode=disable", c.DBHost, c.DBPort, c.DBUser, c.DBPassword, c.DBName)
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	database.AutoMigrate(&Election{})
	database.AutoMigrate(&ValidVoter{})
	database.AutoMigrate(&Candidate{})
	database.AutoMigrate(&CastedVote{})
	database.AutoMigrate(&Vote{})

	db = database
	return nil
}

func TakeDB() *gorm.DB {
	mutex.Lock()
	return db
}

func ReleaseDB() {
	mutex.Unlock()
}

func QueryAllVoters() ([]models.Voter, error) {
	mutex.Lock()
	defer mutex.Unlock()

	// query := `SELECT * FROM valid_voters`
	// rows, err := db.Query(query)
	// if err != nil {
	// 	return nil, err
	// }
	// defer rows.Close()
	// var res []models.Voter

	// for rows.Next() {
	// 	var username models.Voter
	// 	err = rows.Scan(&username)
	// 	res = append(res, username)
	// }

	// return res, nil
	return nil, nil
}

// InsertVoters inserts provided voters into the database, but first queries the database
// to check if they are already added, in which case they are skipped
func InsertVoters(voters []models.Voter) error {
	mutex.Lock()
	defer mutex.Unlock()

	// db.create()

	// query := `SELECT * FROM valid_voters`
	// rows, err := db.Query(query)
	// if err != nil {
	// 	return err
	// }

	// alreadyAdded := map[models.Voter]bool{}
	// for rows.Next() {
	// 	var voter models.Voter
	// 	if err := rows.Scan(&voter); err != nil {
	// 		return err
	// 	}
	// 	alreadyAdded[voter] = true
	// }

	// for _, voter := range voters {
	// 	if !alreadyAdded[voter] {
	// 		if _, err = db.Exec(`INSERT INTO valid_voters VALUES ($1)`, voter); err != nil {
	// 			return err
	// 		}
	// 	}
	// }

	return nil
}

func DeleteVoters(voters []models.Voter) error {
	mutex.Lock()
	defer mutex.Unlock()

	// for _, voter := range voters {
	// 	query := `DELETE FROM valid_voters WHERE email = $1`
	// 	if _, err := db.Exec(query, voter); err != nil {
	// 		return err
	// 	}
	// }

	return nil
}

func InsertElection(e models.Election) error {
	mutex.Lock()
	defer mutex.Unlock()

	// query := `INSERT INTO Elections(id, name, published, finalized, opentime, closetime) VALUES ($1, $2, $3, $4, $5, $6)`
	// _, err := db.Exec(query, e.Id, e.Name, e.IsOpen, e.IsFinalized, e.OpenTime, e.CloseTime)
	// if err != nil {
	// 	println(err)
	// 	return errors.New("Failure while inserting into Elections, see logs for more info")
	// }
	return nil
}

func InsertCandidate(e models.Candidate) error {
	mutex.Lock()
	defer mutex.Unlock()

	// query := `INSERT INTO Candidates(id, name, presentation) VALUES ($1, $2, $3)`
	// if _, err := db.Exec(query, e.Id, e.Name, e.Presentation); err != nil {
	// 	println(err)
	// 	return errors.New("Failure while inserting into Candidates, see logs for more info")
	// }
	return nil
}
