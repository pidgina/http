package middleware

import (
	"log"
	"net/http"
	"time"
)

type responseWriteSpy struct {
	http.ResponseWriter
	status int
}

func (w *responseWriteSpy) WriteHeader(statusCode int) {
	w.status = statusCode
	w.ResponseWriter.WriteHeader((statusCode))
}

func Logging(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		timeNow := time.Now()

		spy := &responseWriteSpy{ResponseWriter: w, status: http.StatusOK}

		log.Printf("Начался запрос с методом: %v, URL: %v", r.Method, r.URL.Path)

		next.ServeHTTP(spy, r)

		if spy.status != http.StatusInternalServerError {
			log.Printf("Запрос с методом: %v, URL: %v Успешно завершен за: %v времени. ", r.Method, r.URL.Path, time.Since(timeNow))
		} else {
			log.Printf("Запрос с методом: %v, URL: %v Поймал ошибку %v, запрос выполнен за %v времени", r.Method, r.URL.Path, spy.status, time.Since(timeNow))

		}

		// log.Printf("Запрос с методом: %v, URL: %v Завершен за: %v времени. Статус: %v", r.Method, r.URL.Path, time.Since(timeNow), spy.status)
	})

}
