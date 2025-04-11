package service

import (
	"log"
	"otus/internal/entity"
	"otus/internal/repository"
)

func GetUser(id int) *entity.User {
	repo, err := repository.NewRepository()
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}

	defer repo.Close()

	user, _ := repo.GetUserByID(id)

	return user
}

func CreateUser(user *entity.User) *entity.User {
	repo, err := repository.NewRepository()
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}

	defer repo.Close()

	createdUser, _ := repo.CreateUser(*user)

	return createdUser
}

func UpdateUser(id int, user *entity.User) *entity.User {
	repo, err := repository.NewRepository()
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}

	defer repo.Close()

	updatedUser, _ := repo.UpdateUser(id, *user)

	return updatedUser
}

func DeleteUser(id int) {
	repo, err := repository.NewRepository()
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}

	defer repo.Close()

	repo.DeleteUserByID(id)
}
