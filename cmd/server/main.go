package main

import (
	"log"
	"net/http"
	"os"
	"proj/internal/handlers"
)

func startServerAndMux() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	mux := http.NewServeMux()

	mux.HandleFunc("GET /time", handlers.TimeNowHandler)
	mux.HandleFunc("GET /time/", handlers.TimeNowHandler)

	mux.HandleFunc("GET /home", handlers.HomePageHandler)
	mux.HandleFunc("GET /home/", handlers.HomePageHandler)

	mux.HandleFunc("GET /hello", handlers.HelloPageHandler)
	mux.HandleFunc("GET /hello/", handlers.HelloPageHandler)

	log.Println("Запуск сервера")
	err := http.ListenAndServe(":"+port, mux)
	if err != nil {
		log.Println("Ошибка запуска сервера:", err)
	}
}

func main() {
	startServerAndMux()
}
