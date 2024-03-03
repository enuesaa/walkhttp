package usecase

import (
	"github.com/enuesaa/walkin/pkg/conf"
	"github.com/enuesaa/walkin/pkg/repository"
)

func CreateConfig(repos repository.Repos) error {
	if err := conf.CreateWalkinDir(repos); err != nil {
		return err
	}

	config := conf.NewConfig()
	if err := repos.Prompt.Ask("BaseUrl", "", &config.BaseUrl); err != nil {
		return err
	}

	return conf.Write(repos, config)
}
