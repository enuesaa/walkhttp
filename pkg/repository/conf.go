package repository

import (
	"encoding/json"
)

type Config struct {
	BaseUrl string `json:"baseUrl"` // example: https://example.com
}

type ConfRepositoryInterface interface {
	NewConfig() Config
	Write(config Config) error
	Read() (Config, error)
	CreateWalkinDir() error
}
type ConfRepository struct{
	fs FsRepository
}

func (repo *ConfRepository) NewConfig() Config {
	return Config{
		BaseUrl: "https://",
	}
}

func (repo *ConfRepository) Write(config Config) error {
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}
	if err := repo.fs.Create(".walkin/config.json", data); err != nil {
		return err
	}

	return nil
}

func (repo *ConfRepository) Read() (Config, error) {
	var conf Config

	data, err := repo.fs.Read(".walkin/config.json")
	if err != nil {
		return conf, err
	}
	if err := json.Unmarshal(data, &conf); err != nil {
		return conf, err
	}

	return conf, nil
}

func (repo *ConfRepository) CreateWalkinDir() error {
	if err := repo.fs.CreateDir(".walkin"); err != nil {
		return err
	}
	return nil
}
