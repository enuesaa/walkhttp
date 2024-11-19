package repository

import "fmt"

func NewDBKeyNotFoundError(key string) error {
	return &DBKeyNotFoundError{
		Key: key,
	}
}

type DBKeyNotFoundError struct {
	Key string
}
func (e *DBKeyNotFoundError) Error() string {
	return fmt.Sprintf("no such key `%s`", e.Key)
}
