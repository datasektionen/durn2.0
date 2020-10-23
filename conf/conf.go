package conf

import (
	"os"
)

type Configuration struct {
	Addr string
}

func ReadConfiguration() Configuration {
	addr := os.Getenv("ADDR")

	return Configuration{Addr: addr}
}
