package usecase

import (
	"github.com/enuesaa/walkin/pkg/conf"
	"github.com/enuesaa/walkin/pkg/repository"
)

func ReadConfig(repos repository.Repos) (conf.Config, error) {
	return conf.Read(repos)
}
