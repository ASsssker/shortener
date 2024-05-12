package main

import (
	"github.com/fatih/color"
	"log"
	"net/http"
	"shortener/cmd/storage"
	"shortener/internal/logger"
)

type Application struct {
	config
	db      *storage.FileDB
	InfoLog *log.Logger
}

func main() {
	var err error

	app := &Application{
		InfoLog: logger.CreateLogger("INFO", color.FgGreen),
	}
	app.parseConfig()

	app.db, err = storage.GetDB(app.config.FileStoragePath)
	if err != nil {
		log.Fatalln(err)
	}
	defer app.db.Close()

	app.InfoLog.Printf("Starting server on %s\n", app.ServerAddr)
	err = http.ListenAndServe(app.ServerAddr, app.getRoutes())
	if err != nil {
		log.Fatalln(err)
	}
}
