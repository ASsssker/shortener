package main

import (
	"context"
	"github.com/fatih/color"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"net/http"
	"shortener/cmd/storage"
	"shortener/internal/logger"
)

type Application struct {
	config
	pgDB     *pgxpool.Pool
	db       *storage.FileDB
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

	app.pgDB, _ = OpenDB(app.DatabaseDSN)
	//if err != nil {
	//	app.ErrorLog.Fatal(err)
	//}
	defer app.pgDB.Close()

	app.db, err = storage.GetDB(app.config.FileStoragePath)
	if err != nil {
		app.ErrorLog.Fatal(err)
	}
	defer app.db.Close()

	app.InfoLog.Printf("Starting server on %s\n", app.ServerAddr)
	err = http.ListenAndServe(app.ServerAddr, app.getRoutes())
	if err != nil {
		app.ErrorLog.Fatal(err)
	}
}

func OpenDB(dataSourceName string) (*pgxpool.Pool, error) {
	dbPool, err := pgxpool.New(context.TODO(), dataSourceName)
	if err != nil {
		return nil, err
	}
	return dbPool, nil
}
