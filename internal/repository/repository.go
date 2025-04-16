package repository

import (
	"database/sql"
	"otus/config"
)

type Repository struct {
	db *sql.DB
}

func NewRepository() (*Repository, error) {
	cfg, err := config.NewConfig()
	db, err := sql.Open("postgres", cfg.PG.URL)
	if err != nil {
		return nil, err
	}

	// Проверяем соединение с базой данных
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &Repository{db: db}, nil
}
