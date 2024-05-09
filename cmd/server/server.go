package main

import (
	"log"
	"net/http"
)

type Application struct {
	config
}

func main() {
	app := &Application{}
	app.parseConfig()

	log.Printf("Starting server on %s\n", app.ServerAddr)
	err := http.ListenAndServe(app.ServerAddr, app.getRoutes())
	if err != nil {
		log.Fatalln(err)
	}
}
