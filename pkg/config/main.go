package config

import (
	"encoding/json"

	"github.com/enuesaa/walkin/pkg/repository"
)

type ConfigFileResource struct {
	Url string `json:"url"`
}

type ConfigFile struct {
	Resources map[string]ConfigFileResource
}

func WriteConfig(repos repository.Repos, config ConfigFile) error {
	fbyte, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}
	return repos.Fs.Create("config.json", fbyte)
}

func ReadConfig(repos repository.Repos) (ConfigFile, error) {
	var config ConfigFile
	fbyte, err := repos.Fs.Read("config.json")
	if err != nil {
		return config, err
	}
	if err := json.Unmarshal(fbyte, &config); err != nil {
		return config, err
	}
	return config, err
}
