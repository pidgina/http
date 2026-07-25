package middleware

import (
	"log"
	"net/http"
	"time"
)

func Logging(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Начался запрос с методом: %v, URL: %v", r.Method, r.URL.Path)
		timeNow := time.Now()
		next.ServeHTTP(w, r)
		timeSince := time.Since(timeNow)
		log.Printf("Запрос с методом: %v, URL: %v Успешно завершен за: %v времени", r.Method, r.URL.Path, timeSince)
	})

}
