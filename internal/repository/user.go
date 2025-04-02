package repository

import (
	"database/sql"
	_ "github.com/lib/pq"
	"otus/internal/entity"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(dataSourceName string) (*Repository, error) {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}

	// Проверяем соединение с базой данных
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &Repository{db: db}, nil
}

func (r *Repository) GetUserByID(id int) (*entity.User, error) {
	var user entity.User
	query := "SELECT id, name FROM users WHERE id = $1"

	row := r.db.QueryRow(query, id)

	if err := row.Scan(&user.ID, &user.Name); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *Repository) Close() error {
	return r.db.Close()
}
