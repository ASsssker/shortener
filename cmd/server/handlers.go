package main

import (
	"io"
	"net/http"
	"shortener/cmd/storage"
)

// PostUrl создает короткий адрес
func (a *Application) PostUrl(w http.ResponseWriter, r *http.Request) {
	url, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	key := generateString(8)
	storage.Urls[key] = string(url)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(key))
}

// GetUrl перенаправляет по адресу
func (a *Application) GetUrl(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	url, ok := storage.Urls[id]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not Found"))
		return
	}

	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}
