package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	config := &Config{}
	config.serverAddr = flag.String("a", ":8080", "server listen address")
	config.rootUrl = flag.String("b", "/", "root url")
	flag.Parse()

	router := getRoutes(*config.rootUrl)
	err := http.ListenAndServe(*config.serverAddr, router)
	if err != nil {
		log.Fatalln(err)
	}
}
