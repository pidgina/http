package handlers

import (
	"fmt"
	"log"
	"net/http"
	"proj/cmd/service"
)

func TimeNowHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte(service.TimeNow()))
	if err != nil {
		log.Println("Ошибка отправки времени в ответ на запрос:", err)
		http.Error(w, "Ошибка отправки времени в ответ на запрос", http.StatusInternalServerError)
		return
	}
}

func HomePageHandler(w http.ResponseWriter, r *http.Request) {

	text := fmt.Sprintln(`Вы находитесь на главной странице, для получения времени используйте страницу:
/time

Для перехода на страницу привествия:
/hello`)

	_, err := w.Write([]byte(text))
	if err != nil {
		log.Println("Ошибка отправки сообщения с домашней страницы:", err)
		http.Error(w, "Ошибка отправки сообщения с домашней страницы", http.StatusInternalServerError)
		return
	}
}

func HelloPageHandler(w http.ResponseWriter, r *http.Request) {
	text := fmt.Sprintln(`Рады приветствовать вас на нашем сайте.
Для перехода на домашнюю страницу со всеми путями используйте:
/home`)

	_, err := w.Write([]byte(text))
	if err != nil {
		log.Println("Ошибка отправки сообщения с приветственной страницы:", err)
		http.Error(w, "Ошибка отправки сообщения с приветственной страницы", http.StatusInternalServerError)
		return
	}

}
