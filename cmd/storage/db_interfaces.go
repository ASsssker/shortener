package storage

import "errors"

type DB interface {
	Get(key string) (string, error)
	Insert(key, value string) error
	Close() error
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
