package main

import (
	"log"
	"net/http"
	"shortener/cmd/handlers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.PostUrl)
	mux.HandleFunc("/{id}", handlers.GetUrl)
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalln(err)
	}
}
