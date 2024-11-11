package repository

import (
	"fmt"
	"maps"
	"slices"
)

type DBRepositoryInterface interface {
	Keys() []string
	Read(key string) (interface{}, error)
	Write(key string, value interface{}) error
}

func NewDBRepository() *DBRepository {
	return &DBRepository{
		data: map[string]interface{}{},
	}
}

type DBRepository struct {
	data map[string]interface{}
}

func (repo *DBRepository) Keys() []string {
	return slices.Sorted(maps.Keys(repo.data))
}

func (repo *DBRepository) Read(key string) (interface{}, error) {
	data, ok := repo.data[key]
	if !ok {
		return nil, fmt.Errorf("no such key")
	}
	return data, nil
}

func (repo *DBRepository) Write(key string, value interface{}) error {
	repo.data[key] = value
	return nil
}
