package database

import (
	"database/sql"
	"errors"
	"fmt"
	"sync"

	"github.com/google/uuid"
	_ "github.com/lib/pq"

	"durn2.0/conf"
	"durn2.0/durn"
)

var mutex sync.Mutex
var db *sql.DB

func CreateDBConnection() {
	c := conf.ReadConfiguration()
	psqlconn := fmt.Sprintf("host=localhost port=%d user=%s password=%s dbname=%s sslmode=disable", c.DBPort, c.DBUser, c.DBPassword, c.DBName)

	mutex.Lock()
	defer mutex.Unlock()
	var err error
	db, err = sql.Open("postgres", psqlconn)
	if err != nil {
		panic(err)
	}
}

func DisconnectDB() {
	mutex.Lock()
	defer mutex.Unlock()

	db.Close()
}

func QueryAllVoters() ([]string, error) {
	mutex.Lock()
	defer mutex.Unlock()

	query := "SELECT * from valid_voters"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var res []string

	for rows.Next() {
		var username string
		err = rows.Scan(&username)
		res = append(res, username)
	}

	return res, nil
}

func InsertElection(durn.Election e) error {
	query := "INSERT INTO Elections(id, name, published, finalized, opentime, closetime) values ($1, $2, $3, $4, $5, $6)"
	_, err := db.Exec(query, e.id, e.name, e.isOpen, e.isFinalized, e.openTime, e.closeTime)
	if err != nil {
		println(err) // TODO: better logging
		return errors.New("Failure while inserting into Elections, see logs for more info")
	}
	return nil
}

func InsertCandidate(durn.Election e) error {
	query := "INSERT INTO Candidates(id, name, presentation) VALUES ($1, $2, $3)"
	_, err := db.Exec(query, e.id, e.name, e.presentation)
	if err != nil {
		println(err) // TODO: better logging
		return errors.New("Failure while inserting into Candidates, see logs for more info")
	}
	return nil
}

func UpdateElectionInfo(durn.Election e) error {
	query := "SELECT * from where id = $1"

	row, err := db.QueryRow(query, e.id)



	query = "UPDATE where id = $1"
	_, err = db.Exec(query, e.id)
}

func