package main

import (
	"time"

	"github.com/rs/zerolog/log"
)

func main() {
	go func() {
		startServer()
		print("\n")
		commands()
	}()
	time.Sleep(1 * time.Second)

	//ctx := context.Background()

	config := getConfig()
	if config.Environment == "development" {
		log.Info().Msgf("Environment: %s", config)
	}
}
