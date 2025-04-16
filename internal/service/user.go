package service

import (
	"otus/internal/container"
	"otus/internal/entity"
)

type UserService struct {
	di container.Container
}

func NewUserService(di container.Container) *UserService {
	return &UserService{di: di}
}

func (userService *UserService) GetUser(id int) *entity.User {
	user, _ := userService.di.Repository.GetUserByID(id)

	return user
}

func (userService *UserService) CreateUser(user *entity.User) *entity.User {

	createdUser, _ := userService.di.Repository.CreateUser(*user)

	return createdUser
}

func (userService *UserService) UpdateUser(id int, user *entity.User) *entity.User {
	updatedUser, _ := userService.di.Repository.UpdateUser(id, *user)

	return updatedUser
}

func (userService *UserService) DeleteUser(id int) {
	userService.di.Repository.DeleteUserByID(id)
}
