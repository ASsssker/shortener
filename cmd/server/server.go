package main

import (
	"log"
	"net/http"
)

func main() {

	router := getRoutes()
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalln(err)
	}
}
