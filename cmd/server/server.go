package main

import (
	"github.com/fatih/color"
	"log"
	"net/http"
	"shortener/cmd/storage"
	db2 "shortener/cmd/storage/db"
	"shortener/cmd/storage/file"
	"shortener/internal/logger"
)

type Application struct {
	config
	DB       storage.DB
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}

func main() {
	var err error
	var db storage.DB

	app := &Application{
		InfoLog:  logger.CreateLogger("INFO", color.FgGreen, log.Ldate|log.Ltime),
		ErrorLog: logger.CreateLogger("ERROR", color.FgRed, log.Ldate|log.Ltime|log.Lshortfile),
	}
	app.parseConfig()

	db, err = db2.NewUrlModel("pgx", app.DatabaseDSN)
	if err != nil {
		app.ErrorLog.Println("Error connecting to database:", err)
		db, err = file.GetDB(app.FileStoragePath)
		if err != nil {
			app.ErrorLog.Println("Failed connect to file storage:", err)
			db = make(storage.Urls)
		}
	}
	app.DB = db

	app.InfoLog.Printf("Starting server on %s\n", app.ServerAddr)
	err = http.ListenAndServe(app.ServerAddr, app.getRoutes())
	if err != nil {
		app.ErrorLog.Fatal(err)
	}
}
