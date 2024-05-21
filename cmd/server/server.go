package main

import (
	"database/sql"
	"github.com/fatih/color"
	_ "github.com/jackc/pgx/v5/pgxpool"
	"log"
	"net/http"
	"shortener/cmd/storage"
	"shortener/internal/logger"
)

type Application struct {
	config
	pgDB     *sql.DB
	FileDB   *storage.FileDB
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}

func main() {
	var err error

	app := &Application{
		InfoLog:  logger.CreateLogger("INFO", color.FgGreen),
		ErrorLog: logger.CreateLogger("ERROR", color.FgRed),
	}
	app.parseConfig()

	app.pgDB, err = OpenDB("pgx", app.DatabaseDSN)
	if err != nil {
		app.FileDB, err = storage.GetDB(app.config.FileStoragePath)
		if err != nil {
			app.ErrorLog.Fatal(err)
		}
		defer app.FileDB.Close()
	}
	defer app.pgDB.Close()

	app.FileDB, err = storage.GetDB(app.config.FileStoragePath)
	if err != nil {
		app.ErrorLog.Fatal(err)
	}
	defer app.FileDB.Close()

	app.InfoLog.Printf("Starting server on %s\n", app.ServerAddr)
	err = http.ListenAndServe(app.ServerAddr, app.getRoutes())
	if err != nil {
		app.ErrorLog.Fatal(err)
	}
}

func OpenDB(driver, dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open(driver, dataSourceName)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
