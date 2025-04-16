package controller

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
	"otus/internal/container"
	"otus/internal/entity"
	"otus/internal/response"
	"otus/internal/service"
	"strconv"
)

func GetUser(di container.Container) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userIdStr := chi.URLParam(r, "id")

		userId, _ := strconv.Atoi(userIdStr)

		userService := service.NewUserService(di)
		user := userService.GetUser(userId)

		if user == nil {
			w.WriteHeader(http.StatusNotFound)
			render.JSON(w, r, response.Error("User not found"))
		} else {
			w.Header().Set("Content-Type", "application/json")
			render.JSON(w, r, user)
		}
	}
}

func CreateUser(di container.Container) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user entity.User
		var createdUser *entity.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, "Не удалось декодировать тело запроса", http.StatusBadRequest)
			return
		}

		userService := service.NewUserService(di)

		createdUser = userService.CreateUser(&user)

		defer r.Body.Close()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(createdUser)
	}
}

func UpdateUser(di container.Container) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		userIdStr := chi.URLParam(r, "id")
		userId, _ := strconv.Atoi(userIdStr)

		userService := service.NewUserService(di)

		user := userService.GetUser(userId)

		if user == nil {
			w.WriteHeader(http.StatusNotFound)
			render.JSON(w, r, response.Error("User not found"))
		} else {
			var requestUser entity.User
			var updatedUser *entity.User
			err := json.NewDecoder(r.Body).Decode(&requestUser)
			if err != nil {
				http.Error(w, "Не удалось декодировать тело запроса", http.StatusBadRequest)
				return
			}

			updatedUser = userService.UpdateUser(userId, &requestUser)

			defer r.Body.Close()

			w.Header().Set("Content-Type", "application/json")
			render.JSON(w, r, updatedUser)
		}
	}
}

func DeleteUser(di container.Container) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userIdStr := chi.URLParam(r, "id")
		userId, _ := strconv.Atoi(userIdStr)

		userService := service.NewUserService(di)

		user := userService.GetUser(userId)
		if user == nil {
			w.WriteHeader(http.StatusNotFound)
			render.JSON(w, r, response.Error("User not found"))
		} else {
			userService.DeleteUser(userId)
			w.Header().Set("Content-Type", "application/json")
			render.JSON(w, r, response.Ok())
		}
	}
}
