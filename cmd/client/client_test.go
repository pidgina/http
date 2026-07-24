package client

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestAPIClient(t *testing.T) {
	NowTime := time.Now()
	layout := "02.01.2006 15:04"
	time := NowTime.Format(layout)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /time", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, time)
	})

	text := fmt.Sprintln(`Вы находитесь на главной странице, для получения времени используйте страницу:
/time

Для перехода на страницу привествия:
/hello`)
	mux.HandleFunc("GET /home", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, text)
	})
	mux.HandleFunc("GET /home/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, text)
	})

	text = fmt.Sprintln(`Рады приветствовать вас на нашем сайте.
Для перехода на домашнюю страницу со всеми путями используйте:
/home`)
	mux.HandleFunc("GET /hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, text)
	})
	mux.HandleFunc("GET /hello/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, text)
	})

	fakeServer := httptest.NewServer(mux)
	defer fakeServer.Close()

	client := NewAPIClient(fakeServer.URL)

	timeResp := client.GetTime()
	homeResp := client.GetHome()
	helloResp := client.GetHello()

	t.Log(timeResp)
	t.Log(homeResp)
	t.Log(helloResp)

}
