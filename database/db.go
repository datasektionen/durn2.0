package database

import (
	"fmt"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"durn2.0/conf"
)

const dsnFormat = "host=%s port=%d user='%s' password='%s' dbname='%s' sslmode=disable"

var mutex sync.Mutex
var db *gorm.DB

// CreateDBConnection Initializes the database connection using
// the connection details specified in env-vars
func CreateDBConnection() error {
	c := conf.ReadConfiguration()
	mutex.Lock()
	defer mutex.Unlock()

	dsn := fmt.Sprintf(dsnFormat, c.DBHost, c.DBPort, c.DBUser, c.DBPassword, c.DBName)
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	err = database.AutoMigrate(&Election{}, &ValidVoter{}, &Candidate{}, &CastedVote{}, &Vote{}, &Ranking{})
	if err != nil {
		return err
	}

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
