package main

import (
	"github.com/go-chi/chi/v5"
)

func (a *Application) getRoutes() *chi.Mux {
	r := chi.NewRouter()
	r.Use(a.RequestsInfo, a.CompressResponse)

	r.Route(a.RootUrl, func(r chi.Router) {
		r.Post("/", a.PostUrl)
		r.Get("/{id}", a.GetUrl)
	})
	return r
}
