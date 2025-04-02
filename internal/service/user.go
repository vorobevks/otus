package service

import (
	"log"
	"otus/internal/entity"
	"otus/internal/repository"
)

func GetUser() *entity.User {
	dataSourceName := "user=otus password=password dbname=otus sslmode=disable port=5435"

	repo, err := repository.NewRepository(dataSourceName)
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}

	defer repo.Close()

	userID := 1 // Пример ID пользователя для поиска
	user, _ := repo.GetUserByID(userID)

	return user
}
