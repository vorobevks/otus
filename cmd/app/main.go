package main

import (
	"log"
	"otus/config"
	"otus/internal/app"
	"otus/internal/container"
	"otus/internal/repository"
	"otus/pkg/logger"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}
	repo, _ := repository.NewRepository()
	l := logger.New(cfg.Log.Level)

	di := container.New(cfg, repo)

	// Run
	app.Run(di, l)
}
