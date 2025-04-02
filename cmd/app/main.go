package main

import (
	"log"
	"otus/config"
	"otus/internal/app"
	"otus/internal/container"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}
	di := container.New(cfg)

	// Run
	app.Run(di)
}
