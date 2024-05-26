package main

import (
	"github.com/fatih/color"
	"log"
	"net/http"
	"shortener/internal/logger"
)

type Application struct {
	config
	DB       storageRepo
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}

func main() {
	var err error
	app := initApp()

	app.InfoLog.Printf("Starting server on %s\n", app.ServerAddr)
	err = http.ListenAndServe(app.ServerAddr, app.getRoutes())
	if err != nil {
		app.ErrorLog.Fatal(err)
	}
}

func initApp() *Application {
	app := &Application{
		InfoLog:  logger.CreateLogger("INFO", color.FgGreen, log.Ldate|log.Ltime, nil),
		ErrorLog: logger.CreateLogger("ERROR", color.FgRed, log.Ldate|log.Ltime|log.Lshortfile, nil),
	}
	app.parseConfig()
	app.connectToStorage()
	
	return app
}
