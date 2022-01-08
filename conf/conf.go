package conf

import (
	"fmt"
	"log"
	"os"
	"strconv"

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
		log.Fatal(fmt.Sprintf("panic: Env var '%s' not set", varName))
	}
	return val
}

func readEnvFallback(varName string, fallback string) string {
	val, present := os.LookupEnv(varName)
	if !present {
		return fallback
	}
	return val
}

func readEnvInteger(varName string, fallback int) int {
	val, present := os.LookupEnv(varName)
	if !present {
		return fallback
	}
	num, err := strconv.Atoi(val)
	if err != nil {
		log.Fatal(fmt.Sprintf("panic: Env var '%s' could not be parsed as integer", varName))
	}
	return num
}

func readEnvBoolean(varName string, fallback bool) bool {
	val, present := os.LookupEnv(varName)
	if !present {
		return fallback
	}
	bool, err := strconv.ParseBool(val)
	if err != nil {
		log.Fatal(fmt.Sprintf("panic: Env var '%s' could not be parsed as boolean", varName))
	}
	return bool
}

var conf Configuration
var initialized = false

func GetConfiguration() Configuration {

	if initialized {
		return conf
	}

	if err := dotenv.Load(); err != nil {
		log.Println("No .env found")
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

	initialized = true
	return conf
}
