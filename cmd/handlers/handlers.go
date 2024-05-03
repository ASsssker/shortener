package handlers

import (
	"io"
	"net/http"
	"shortener/cmd/storage"
)

// PostUrl создает короткий адрес
func PostUrl(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method Not Allowed"))
		return
	}

	url, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	key := GenerateString(8)
	storage.Urls[key] = string(url)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(key))
}

// GetUrl перенаправляет по адресу
func GetUrl(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method Not Allowed"))
		return
	}

	id := r.PathValue("id")
	url, ok := storage.Urls[id]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not Found"))
		return
	}

	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}
