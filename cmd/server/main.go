package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"proj/internal/handlers"
	"proj/internal/middleware"
	"syscall"
	"time"
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

	res := middleware.Logging(mux)

	server := &http.Server{Addr: ":" + port, Handler: res}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Ошибка сервера: %v", err)
		}
	}()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	<-sigCh

	fmt.Println("Получен сигнал завершения, останавливаем сервер...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Ошибка при остановке сервера: %v", err)
	}

	fmt.Println("Сервер остановлен корректно")
}

func main() {
	fmt.Println("Запускаем сервер.")
	startServerAndMux()
}
