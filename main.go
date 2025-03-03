package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Message string `json:"status"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := Response{Message: "ok"}
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/health", handler)
	fmt.Println("Сервер запущен на порту 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Ошибка при запуске сервера:", err)
	}
}
