package usecase

import (
	"encoding/json"

	"github.com/enuesaa/walkin/pkg/repository"
)

type WalkinConf struct {
	BaseUrl string `json:"baseUrl"`
}
func CreateConfig(repos repository.Repos) error {
	if err := repos.Fs.CreateDir(".walkin"); err != nil {
		return err
	}

	conf := WalkinConf{
		BaseUrl: "https://",
	}
	if err := repos.Prompt.Ask("BaseUrl", "", &conf.BaseUrl); err != nil {
		return err
	}

	data, err := json.Marshal(conf)
	if err != nil {
		return err
	}
	if err := repos.Fs.Create(".walkin/config.json", data); err != nil {
		return err
	}

	return nil
}
