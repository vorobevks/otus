package config

import (
	"fmt"
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type (
	// Config -.
	Config struct {
		Log Log
		PG  PG
	}

	// Log -.
	Log struct {
		Level string `env:"LOG_LEVEL,required"`
	}

	PG struct {
		PoolMax int    `env:"PG_POOL_MAX,required"`
		URL     string `env:"PG_URL,required"`
	}
)

// NewConfig returns app config.
func NewConfig() (*Config, error) {
	cfg := &Config{}

	if err := godotenv.Load(".env"); err != nil {
		panic("Ошибка загрузки .env файла: " + err.Error())
	}

	if err := env.Parse(cfg); err != nil {
		return cfg, fmt.Errorf("consdfsdfsdffig error: %w", err)
	}

	return cfg, nil
}
