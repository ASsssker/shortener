package db

import (
	"database/sql"
	"errors"
	"shortener/internal/storage"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type Url struct {
	ID      int
	Url     string
	Key     string
	Created time.Time
}

type UrlModel struct {
	DB *sql.DB
}

func NewUrlModel(driver, dataSourceName string) (*UrlModel, error) {
	db, err := OpenDB(driver, dataSourceName)
	if err != nil {
		return nil, err
	}

	u := &UrlModel{DB: db}
	err = u.initTable()
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (m *UrlModel) initTable() error {
	stmt := "CREATE TABLE IF NOT EXISTS urls (id SERIAL PRIMARY KEY, keys TEXT NOT NULL, url TEXT NOT NULL, created TIMESTAMP NOT NULL)"
	stmtCreateIdx := "CREATE INDEX IF NOT EXISTS idx_url ON urls(url)"
	if _, err := m.DB.Exec(stmt); err != nil {
		return err
	}
	if _, err := m.DB.Exec(stmtCreateIdx); err != nil {
		return err
	}
	return nil
}

func (m *UrlModel) Get(key string) (string, error) {
	stmt := `SELECT id, url, keys, created FROM urls WHERE keys = $1`

	u := &Url{}
	err := m.DB.QueryRow(stmt, key).Scan(&u.ID, &u.Url, &u.Key, &u.Created)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", storage.ErrNoRecord
		}
		return "", err
	}
	return u.Url, nil
}

func (m *UrlModel) Insert(key string, url string) (string, string, bool ,error) {
	stmt := `INSERT INTO urls (keys, url, created) VALUES ($1, $2, $3)`

	u, exists := m.checkExists(url)
	if !exists {
		_, err := m.DB.Exec(stmt, key, url, time.Now())
		if err != nil {
			return "", "", false, err
		}
	} else {
		key = u.Key
	}
	return key, url, exists, nil
}

func (m *UrlModel) checkExists(url string) (*Url, bool) {
	stmt := "SELECT id, url, keys, created FROM urls WHERE url = $1"

	u := &Url{}
	err := m.DB.QueryRow(stmt, url).Scan(&u.ID, &u.Url, &u.Key, &u.Created)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return u, false
		}
	}
	return u, true
}

func (m *UrlModel) Close() error {
	return m.DB.Close()
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
