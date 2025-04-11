package service

import (
	"crypto/md5"
	"encoding/hex"
	"log"
	"otus/internal/entity"
	"otus/internal/repository"
)

func GetUser(id int) *entity.User {
	dataSourceName := "user=otus password=password dbname=otus sslmode=disable port=5435"

	repo, err := repository.NewRepository(dataSourceName)
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}

	defer repo.Close()

	user, _ := repo.GetUserByID(id)

	return user
}

func CreateUser(user *entity.User) *entity.User {
	dataSourceName := "user=otus password=password dbname=otus sslmode=disable port=5435"
	repo, err := repository.NewRepository(dataSourceName)
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}

	defer repo.Close()

	user.Password = getMD5Hash(user.Password)

	createdUser, _ := repo.CreateUser(*user)

	return createdUser
}

func getMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
