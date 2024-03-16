package usecase

import (
	"github.com/enuesaa/walkin/pkg/repository"
)

func CreateConfig(repos repository.Repos) error {
	if err := repos.Conf.CreateWalkinDir(); err != nil {
		return err
	}

	config := repos.Conf.NewConfig()
	if err := repos.Prompt.Ask("BaseUrl", "", &config.BaseUrl); err != nil {
		return err
	}

	return repos.Conf.Write(config)
}
