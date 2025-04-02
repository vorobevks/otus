package controller

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"otus/internal/service"
	"strconv"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	userIdStr := chi.URLParam(r, "id")

	userId, _ := strconv.Atoi(userIdStr)

	user := service.GetUser(userId)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
