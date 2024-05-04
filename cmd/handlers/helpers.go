package handlers

import (
	"bytes"
	"github.com/go-resty/resty/v2"
	"math/rand"
	"net/http"
	"net/http/httptest"
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

// testRequest делает запрос хэндлеру через сервер
func testRequest(srv *httptest.Server, method string) (*resty.Response, error) {
	req := resty.New().R()
	req.Method = method
	req.URL = srv.URL
	resp, err := req.Send()
	return resp, err
}
