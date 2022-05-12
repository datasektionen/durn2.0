package conf

import (
	"os"
)

type Configuration struct {
	Addr string
	LoginApiKey string
}

func ReadConfiguration() Configuration {
	addr := os.Getenv("ADDR")
	loginApiKey := os.Getenv("LOGIN_API_KEY")

	return Configuration{
		Addr:        addr,
		LoginApiKey: loginApiKey,
	}
}
