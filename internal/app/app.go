package app

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
	"otus/internal/container"
	"otus/internal/controller"
	"otus/internal/response"
	"otus/pkg/logger"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	resp := response.Ok()
	json.NewEncoder(w).Encode(resp)
}

func Run(di container.Container, logger *logger.Logger) {
	r := chi.NewRouter()
	r.Get("/user/{id}", controller.GetUser(di))
	r.Post("/user", controller.CreateUser(di))
	r.Put("/user/{id}", controller.UpdateUser(di))
	r.Delete("/user/{id}", controller.DeleteUser(di))
	r.Get("/health", handler)

	fmt.Println("Сервер запущен на порту 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Println("Ошибка при запуске сервера:", err)
	}
}
