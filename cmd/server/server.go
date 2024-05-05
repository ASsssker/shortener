package main

import (
	"log"
	"net/http"
)

func main() {
	cfg, err := getConfig()
	if err != nil {
		log.Fatal(err)
	}

	router := getRoutes(cfg.RootUrl)
	log.Printf("Starting server on %s\n", cfg.ServerAddr)
	err = http.ListenAndServe(cfg.ServerAddr, router)
	if err != nil {
		log.Fatalln(err)
	}
}
