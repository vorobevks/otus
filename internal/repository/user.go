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
	query := "SELECT id, username, name, created_at  FROM users WHERE id = $1"

	row := r.db.QueryRow(query, id)

	if err := row.Scan(
		&user.ID,
		&user.Username,
		&user.Name,
		&user.CreatedAt,
	); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *Repository) CreateUser(user entity.User) (*entity.User, error) {
	var id int
	var createdAt string
	query := "INSERT INTO users (username, name, password) VALUES ($1, $2, $3) RETURNING id, created_at"
	err := r.db.QueryRow(query, user.Username, user.Name, user.Password).Scan(&id, &createdAt)
	if err != nil {
		return &entity.User{}, err
	}

	return &entity.User{
		ID:        id,
		Username:  user.Username,
		Name:      user.Name,
		CreatedAt: createdAt,
	}, nil
}

func (r *Repository) Close() error {
	return r.db.Close()
}
