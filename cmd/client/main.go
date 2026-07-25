package main

import (
	"fmt"
	"os"
	"proj/internal/api"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println(api.NewAPIClient("http://localhost:" + port).GetHello())
	fmt.Println(api.NewAPIClient("http://localhost:" + port).GetHome())
	fmt.Println(api.NewAPIClient("http://localhost:" + port).GetTime())

}
