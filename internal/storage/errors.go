package storage

import "errors"

var (
	ErrNoRecord = errors.New("no matching record")
	ErrRecordExistings = errors.New("record already exists")
)