package storage

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
)

type Urls map[string]string

type FileDB struct {
	file *os.File // Файл для хранения данных. Явлется опциональным
	Urls          // Мапа для хранения данных
}

// GetDB инициализирует структуру для хранения данных.
func GetDB(dbFilePath string) (*FileDB, error) {
	db := &FileDB{Urls: make(Urls)}

	if dbFilePath != "" {
		f, err := os.OpenFile(dbFilePath, os.O_RDWR|os.O_CREATE, 0774)
		if err != nil {
			return nil, err
		}

		db.file = f
		if err := json.NewDecoder(f).Decode(&db.Urls); err != nil && err != io.EOF { //Если файл пуст игнорируем ошибку
			return nil, err
		}
	}
	return db, nil
}

func (f *FileDB) Close() error {
	return f.file.Close()
}

// UpdateFile синхронизирует данные в мапе и файле
func (f *FileDB) UpdateFile() error {
	if f.file != nil {
		var buf bytes.Buffer
		if err := json.NewEncoder(&buf).Encode(f.Urls); err != nil {
			return err
		}
		if err := f.file.Truncate(0); err != nil {
			return err
		}
		if _, err := f.file.Write(buf.Bytes()); err != nil {
			return err
		}
	}
	return nil
}