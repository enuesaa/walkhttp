package repository

import (
	"fmt"
	"sort"
)

type DBRepositoryInterface interface {
	List() []string
	Get(id string) (interface{}, error)
	Save(id string, data interface{}) error
}

type InmemoryDB struct{
	data map[string]interface{}
}

func (repo *InmemoryDB) List() []string {
	keys := make([]string, 0)
    for k := range repo.data {
        keys = append(keys, k)
    }
	sort.Strings(keys)

	return keys
}

func (repo *InmemoryDB) Get(id string) (interface{}, error) {
	data, ok := repo.data[id]
	if !ok {
		return nil, fmt.Errorf("not found")
	}
	return data, nil
}

func (repo *InmemoryDB) Save(id string, data interface{}) error {
	repo.data[id] = data
	return nil
}
