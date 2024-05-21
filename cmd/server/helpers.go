package main

import (
	"bytes"
	"github.com/go-resty/resty/v2"
	"log"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"shortener/cmd/storage"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// generateString генерирует случайную строку указанной длины
func generateString(length int) string {
	s := make([]rune, length)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

// runHandler делает запрос к хэндлеру без запуска сервера
func runHandler(method string, url string, body string, f func(http.ResponseWriter, *http.Request)) *http.Response {
	r := httptest.NewRequest(method, url, bytes.NewReader([]byte(body)))
	w := httptest.NewRecorder()
	f(w, r)
	return w.Result()
}

// testRequest делает запрос к хэндлеру через сервер
func testRequest(srv *httptest.Server, method string, body string) (*resty.Response, error) {
	req := resty.New().R()
	req.Method = method
	req.URL = srv.URL
	req.Body = body
	resp, err := req.Send()
	return resp, err
}

func getTestApp() *Application {
	db, err := storage.GetDB("")
	if err != nil {
		log.Fatal(err)
	}

	return &Application{FileDB: db}
}
