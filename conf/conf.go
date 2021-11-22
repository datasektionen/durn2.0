package conf

import (
	"os"
	"strconv"
)

type Configuration struct {
	Addr        string
	LoginApiKey string
	DBPort      int
	DBUser      string
	DBPassword  string
	DBName      string
}

func ReadConfiguration() Configuration {
	addr := os.Getenv("ADDR")
	loginApiKey := os.Getenv("LOGIN_API_KEY")

	port := os.Getenv("DB_PORT")
	dbPort := 5432
	if port != "" {
		dbPort, _ = strconv.Atoi(port)
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	return Configuration{
		Addr:        addr,
		LoginApiKey: loginApiKey,
		DBPort:      dbPort,
		DBUser:      dbUser,
		DBPassword:  dbPassword,
		DBName:      dbName,
	}
}
