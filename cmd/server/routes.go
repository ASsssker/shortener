package main

import (
	"github.com/go-chi/chi/v5"
	"shortener/cmd/handlers"
)

func getRoutes() *chi.Mux {
	r := chi.NewRouter()

	r.Post("/", handlers.PostUrl)
	r.Get("/{id}", handlers.GetUrl)

	return r
}
