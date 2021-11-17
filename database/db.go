package database

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/lib/pq"

	"durn2.0/conf"
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

func QueryAllVoters() []string {
	mutex.Lock()
	defer mutex.Unlock()

	query := "SELECT * from valid_voters"
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var res []string

	for rows.Next() {
		var username string
		err = rows.Scan(&username)
		res = append(res, username)
	}

	return res
}
