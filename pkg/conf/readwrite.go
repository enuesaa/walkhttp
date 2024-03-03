package conf

import (
	"encoding/json"

	"github.com/enuesaa/walkin/pkg/repository"
)

func Write(repos repository.Repos, config Config) error {
	data, err := json.Marshal(config)
	if err != nil {
		return err
	}
	if err := repos.Fs.Create(".walkin/config.json", data); err != nil {
		return err
	}

	return nil
}

func Read(repos repository.Repos) (Config, error) {
	var conf Config

	data, err := repos.Fs.Read(".walkin/config.json")
	if err != nil {
		return conf, err
	}
	if err := json.Unmarshal(data, &conf); err != nil {
		return conf, err
	}

	return conf, nil
}
