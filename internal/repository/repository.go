package repository

import (
	"database/sql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
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

	driver, _ := postgres.WithInstance(db, &postgres.Config{})
	m, _ := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres", driver)
	m.Up()

	return &Repository{db: db}, nil
}
