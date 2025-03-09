package repository

import (
	"maps"
	"slices"
)

type DBRepositoryInterface interface {
	Keys() []string
	Read(key string) (any, error)
	Write(key string, value any) error
	Omit(key string) error
	Count() int
}

type DBRepository struct {
	data map[string]any
}

func (repo *DBRepository) Keys() []string {
	return slices.Sorted(maps.Keys(repo.data))
}

func (repo *DBRepository) Read(key string) (any, error) {
	data, ok := repo.data[key]
	if !ok {
		return nil, ErrDBKeyNotFound
	}
	return data, nil
}

func (repo *DBRepository) Write(key string, value any) error {
	repo.data[key] = value
	return nil
}

func (repo *DBRepository) Omit(key string) error {
	_, ok := repo.data[key]
	if !ok {
		return ErrDBKeyNotFound
	}

	delete(repo.data, key)
	return nil
}

func (repo *DBRepository) Count() int {
	return len(repo.data)
}
