package usecase

import (
	"encoding/json"

	"github.com/enuesaa/walkin/pkg/repository"
)

func LoadConfig(repos repository.Repos) (WalkinConf, error) {
	var conf WalkinConf

	data, err := repos.Fs.Read(".walkin/config.json")
	if err != nil {
		return conf, err
	}
	if err := json.Unmarshal(data, &conf); err != nil {
		return conf, err
	}

	return conf, nil
}
