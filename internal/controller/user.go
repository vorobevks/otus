package controller

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"otus/internal/entity"
	"otus/internal/service"
	"strconv"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	userIdStr := chi.URLParam(r, "id")

	userId, _ := strconv.Atoi(userIdStr)

	user := service.GetUser(userId)

	if user == nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("user not found")
	} else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	}
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user entity.User
	var createdUser *entity.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Не удалось декодировать тело запроса", http.StatusBadRequest)
		return
	}

	createdUser = service.CreateUser(&user)

	defer r.Body.Close()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdUser)
}
