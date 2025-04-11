package app

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
	"otus/internal/container"
	"otus/internal/controller"
	"otus/pkg/logger"
	"otus/pkg/postgres"
)

type Response struct {
	Message string `json:"status"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := Response{Message: "ok"}
	json.NewEncoder(w).Encode(response)
}

func Run(di container.Container) {
	l := logger.New(di.Config.Log.Level)

	pg, err := postgres.New(di.Config.PG.URL, postgres.MaxPoolSize(di.Config.PG.PoolMax))

	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	defer pg.Close()

	r := chi.NewRouter()
	r.Get("/user/{id}", controller.GetUser)
	r.Post("/user", controller.CreateUser)
	r.Put("/user/{id}", controller.UpdateUser)
	r.Get("/health", handler)

	fmt.Println("Сервер запущен на порту 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Println("Ошибка при запуске сервера:", err)
	}
}
