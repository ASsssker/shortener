package main

import (
	"errors"
	"shortener/internal/storage/db"
	"shortener/internal/storage/file"
)

type storageRepo interface {
	Get(key string) (string, error)
	Insert(key, value string) error
	Close() error
}

// ConnectToStorage подключает к хранилищу со следующим приоритетом: SQL > file > RAM
func (app *Application) connectToStorage() {
	var storage storageRepo
	var err error
	if storage, err = db.NewUrlModel("pgx", app.DatabaseDSN); err == nil {
		app.DB = storage
		app.InfoLog.Print("Connected to Postgres")
		return
	}
	app.ErrorLog.Print(err)

	if storage, err = file.GetDB(app.FileStoragePath); err == nil {
		app.DB = storage
		app.InfoLog.Print("Connected to file storage")
		return
	}
	app.ErrorLog.Print(err)

	storage = make(Urls)
	app.DB = storage

	app.InfoLog.Print("Connected to RAM storage")
}


type Urls map[string]string

func (u Urls) Get(key string) (string, error) {
	value, exists := u[key]
	if !exists {
		return "", errors.New("url not found")
	}
	return value, nil
}

func (u Urls) Insert(key, value string) error {
	_, exists := u[key]
	if exists {
		return errors.New("url already exists")
	}
	return nil
}

func (u Urls) Close() error {
	return nil
}
