package server

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestTimeNow(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/time", nil)

	rec := httptest.NewRecorder()

	TimeNowHandler(rec, req)

	if req.Method != http.MethodGet {
		t.Fatalf("method = %v, want %v", req.Method, http.MethodGet)
	}

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}

	NowTime := time.Now()
	layout := "02.01.2006 15:04"
	time := NowTime.Format(layout)

	if rec.Body.String() != time {
		t.Fatalf("body = %v, want %v", rec.Body.String(), time)
	}

}

func TestHomePage(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/home", nil)

	rec := httptest.NewRecorder()

	if req.Method != http.MethodGet {
		t.Fatalf("method = %v, want %v", req.Method, http.MethodGet)
	}

	HomePageHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}

	text := fmt.Sprintln(`Вы находитесь на главной странице, для получения времени используйте страницу:
/time

Для перехода на страницу привествия:
/hello`)

	if rec.Body.String() != text {
		t.Fatalf("body = %v, want %v", rec.Body.String(), text)
	}

}

func TestHelloPage(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/hello", nil)

	rec := httptest.NewRecorder()

	if req.Method != http.MethodGet {
		t.Fatalf("method = %v, want %v", req.Method, http.MethodGet)
	}

	HelloPageHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %v, want %v", rec.Code, http.StatusOK)
	}

	text := fmt.Sprintln(`Рады приветствовать вас на нашем сайте.
Для перехода на домашнюю страницу со всеми путями используйте:
/home`)

	if rec.Body.String() != text {
		t.Fatalf("body = %v, want %v", rec.Body.String(), text)
	}

}
