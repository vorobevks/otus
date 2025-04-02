package controller

import (
	"encoding/json"
	"net/http"
	"otus/internal/entity"
	"otus/internal/service"
)

type Response struct {
	User *entity.User
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	user := service.GetUser()

	w.Header().Set("Content-Type", "application/json")
	response := Response{User: user}
	json.NewEncoder(w).Encode(response)
}
