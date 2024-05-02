package main

import (
	"log"
	"net/http"
)

var urls = make(map[string]string)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", PostUrl)
	mux.HandleFunc("/{id}", GetUrl)
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalln(err)
	}
}
