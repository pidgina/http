package main

import (
	// "proj/cmd/client"
	"fmt"
	"log"
	"net/http"
	"proj/cmd/client"
	"proj/cmd/server"
	"sync"
	"time"
)

var wg sync.WaitGroup

func startServerAndMux() {
	defer wg.Done()
	mux := http.NewServeMux()

	mux.HandleFunc("GET /time", server.TimeNowHandler)
	mux.HandleFunc("GET /time/", server.TimeNowHandler)

	mux.HandleFunc("GET /home", server.HomePageHandler)
	mux.HandleFunc("GET /home/", server.HomePageHandler)

	mux.HandleFunc("GET /hello", server.HelloPageHandler)
	mux.HandleFunc("GET /hello/", server.HelloPageHandler)

	log.Println("Запуск сервера")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Println("Ошибка запуска сервера:", err)
	}

}

func startClient() {
	client := client.NewAPIClient("http://localhost:8080")
	fmt.Println(client.GetTime())
	fmt.Println(client.GetHome())
	fmt.Println(client.GetHello())
}

func main() {
	// wg.Add(1)

	// for i := 0; i < 1; i++ {
	// 	wg.Add(1)
	// 	go startServerAndMux()
	// }

	// wg.Wait()
	wg.Add((1))
	go startServerAndMux()

	// startServerAndMux()
	for i := 0; i < 5; i++ {
		time.Sleep(2 * time.Second)
		log.Println("Отправка номер:", i+1)
		startClient()

	}
	log.Println("Произвели 5 итераций отправки.")

	wg.Wait()

}
