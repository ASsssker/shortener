package handlers

import (
	"bytes"
	"math/rand"
	"net/http"
	"net/http/httptest"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func generateString(length int) string {
	s := make([]rune, length)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func request(method string, url string, body string, f func(http.ResponseWriter, *http.Request)) *http.Response {
	r := httptest.NewRequest(method, url, bytes.NewReader([]byte(body)))
	w := httptest.NewRecorder()
	f(w, r)
	return w.Result()
}
