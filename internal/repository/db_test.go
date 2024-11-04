package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDBRepositoryKeys(t *testing.T) {
	repo := DBRepository{
		data: map[string]interface{}{},
	}

	repo.Write("key2", "value2")
	repo.Write("key1", "value1")
	repo.Write("key3", "value3")

	keys := repo.Keys()
	assert.Equal(t, []string{"key1", "key2", "key3"}, keys)
}

func TestDBRepositoryReadWrite(t *testing.T) {
	repo := DBRepository{
		data: map[string]interface{}{},
	}

	err := repo.Write("key1", "value1")
	assert.NoError(t, err)

	err = repo.Write("key2", "value2")
	assert.NoError(t, err)

	val, err := repo.Read("key1")
	assert.NoError(t, err)
	assert.Equal(t, "value1", val)

	val, err = repo.Read("key2")
	assert.NoError(t, err)
	assert.Equal(t, "value2", val)

	_, err = repo.Read("non_existent_key")
	assert.Error(t, err)
	assert.EqualError(t, err, "no such key")
}
