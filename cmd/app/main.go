package main

import (
	"bukus/config"
	"bukus/internal/app"
	"log"
)

func main() {
	// Configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg)
}
