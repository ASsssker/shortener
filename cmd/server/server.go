package main

import (
	"github.com/fatih/color"
	"log"
	"net/http"
	"shortener/internal/logger"
)

type Application struct {
	config
	InfoLog *log.Logger
}

func main() {
	app := &Application{
		InfoLog: logger.CreateLogger("INFO", color.FgGreen),
	}
	app.parseConfig()

	app.InfoLog.Printf("Starting server on %s\n", app.ServerAddr)
	err := http.ListenAndServe(app.ServerAddr, app.getRoutes())
	if err != nil {
		log.Fatalln(err)
	}
}
