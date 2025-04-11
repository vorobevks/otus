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
	query := "SELECT id, username, first_name, last_name, email, phone  FROM users WHERE id = $1"

	row := r.db.QueryRow(query, id)

	if err := row.Scan(
		&user.ID,
		&user.Username,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Phone,
	); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *Repository) CreateUser(user entity.User) (*entity.User, error) {
	var id int
	query := "INSERT INTO users (username, first_name, last_name, email, phone) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	err := r.db.QueryRow(query, user.Username, user.FirstName, user.LastName, user.Email, user.Phone).Scan(&id)
	if err != nil {
		return &entity.User{}, err
	}

	return &entity.User{
		ID:        id,
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Phone:     user.Phone,
	}, nil
}

func (r *Repository) UpdateUser(id int, user entity.User) (*entity.User, error) {
	var username string
	query := "UPDATE users SET first_name = $1, last_name = $2, email = $3, phone = $4 WHERE id = $5 RETURNING id, username"
	err := r.db.QueryRow(query, user.FirstName, user.LastName, user.Email, user.Phone, id).Scan(&id, &username)
	if err != nil {
		return &entity.User{}, err
	}

	return &entity.User{
		ID:        id,
		Username:  username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Phone:     user.Phone,
	}, nil
}

func (r *Repository) Close() error {
	return r.db.Close()
}
