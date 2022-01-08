package conf

import (
	"context"
	"fmt"
	"os"
	"strconv"

	rl "durn2.0/requestLog"
	dotenv "github.com/joho/godotenv"
)

type Configuration struct {
	Addr        string
	LoginApiKey string
	DBPort      int
	DBUser      string
	DBPassword  string
	DBName      string
	DBHost      string
	SkipAuth    bool
}

func readEnvRequired(varName string) string {
	val, precent := os.LookupEnv(varName)
	if !precent {
		rl.Fatal(context.Background(), fmt.Sprintf("panic: Env var '%s' not set", varName))
		panic("exiting")
	}
	return val
}

func readEnvFallback(varName string, fallback string) string {
	val, precent := os.LookupEnv(varName)
	if !precent {
		return fallback
	}
	return val
}

func readEnvInteger(varName string, fallback int) int {
	val, precent := os.LookupEnv(varName)
	if !precent {
		return fallback
	}
	num, err := strconv.Atoi(val)
	if err != nil {
		rl.Fatal(context.Background(), fmt.Sprintf("panic: Env var '%s' is not an integer", varName))
		panic("exiting")
	}
	return num
}

func readEnvBoolean(varName string, fallback bool) bool {
	val, precent := os.LookupEnv(varName)
	if !precent {
		return fallback
	}
	return val == "true"
}

var conf Configuration
var intialized bool = false

func GetConfiguration() Configuration {

	if intialized {
		return conf
	}

	if err := dotenv.Load(); err != nil {
		rl.Info(context.Background(), "No .env found")
	}

	conf = Configuration{
		Addr:        readEnvFallback("ADDR", "localhost:8080"),
		LoginApiKey: readEnvRequired("LOGIN_API_KEY"),
		DBHost:      readEnvFallback("DB_HOST", "localhost"),
		DBPort:      readEnvInteger("DB_PORT", 5432),
		DBUser:      readEnvRequired("DB_USER"),
		DBPassword:  readEnvRequired("DB_PASSWORD"),
		DBName:      readEnvRequired("DB_NAME"),
		SkipAuth:    readEnvBoolean("SKIP_AUTH", false),
	}

	intialized = true
	return conf
}
