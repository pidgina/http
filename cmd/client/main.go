package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type APIClient struct {
	baseURL string
	http    *http.Client
}

func NewAPIClient(url string) *APIClient {
	return &APIClient{
		baseURL: url,
		http:    &http.Client{Timeout: 5 * time.Second},
	}
}

func (c *APIClient) GetTime() string {
	resp, err := c.http.Get(c.baseURL + "/time")
	if err != nil {
		log.Println("При отправки на сервер GET запроса по адресу /time произошла ошибка:", err)
		return ""
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println("Неверный путь или ошибка сервера. Статус:", resp.StatusCode)
		return ""
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Ошибка чтения тела ответа:", err)
		return ""
	}
	return string(data)
}

func (c *APIClient) GetHome() string {
	resp, err := c.http.Get(c.baseURL + "/home")
	if err != nil {
		log.Println("При отправки на сервер GET запроса по адресу /home произошла ошибка:", err)
		return ""
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println("Неверный путь или ошибка сервера. Статус:", resp.StatusCode)
		return ""
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Ошибка чтения тела ответа:", err)
		return ""
	}
	return string(data)
}

func (c *APIClient) GetHello() string {
	resp, err := c.http.Get(c.baseURL + "/hello")
	if err != nil {
		log.Println("При отправки на сервер GET запроса по адресу /hello произошла ошибка:", err)
		return ""
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println("Неверный путь или ошибка сервера. Статус:", resp.StatusCode)
		return ""
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Ошибка чтения тела ответа:", err)
		return ""
	}
	return string(data)
}

func main() {
	client := NewAPIClient("http://localhost:8080")
	fmt.Println(client.GetTime())
	fmt.Println(client.GetHome())
	fmt.Println(client.GetHello())

}
