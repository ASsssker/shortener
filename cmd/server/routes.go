package main

import (
	"github.com/go-chi/chi/v5"
	"shortener/cmd/handlers"
)

func getRoutes(rootUrl string) *chi.Mux {
	r := chi.NewRouter()
	r.Route(rootUrl, func(r chi.Router) {
		r.Post("/", handlers.PostUrl)
		r.Get("/{id}", handlers.GetUrl)
	})

	return r
}
