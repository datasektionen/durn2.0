package conf

import (
	"fmt"
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
}

func readEnvRequired(varName string) string {
	val, precent := os.LookupEnv(varName)
	if !precent {
		panic(fmt.Sprintf("panic: Env var '%s' not set", varName))
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
		panic(fmt.Sprintf("panic: Env var '%s' is not an integer", varName))
	}
	return num
}

func ReadConfiguration() Configuration {
	if err := dotenv.Load(); err != nil {
		fmt.Println("No .env found")
	}

	return Configuration{
		Addr:        readEnvFallback("ADDR", "localhost"),
		LoginApiKey: readEnvRequired("LOGIN_API_KEY"),
		DBHost:      readEnvFallback("DB_HOST", "localhost"),
		DBPort:      readEnvInteger("DB_PORT", 5432),
		DBUser:      readEnvRequired("DB_USER"),
		DBPassword:  readEnvRequired("DB_PASSWORD"),
		DBName:      readEnvRequired("DB_NAME"),
	}
}
