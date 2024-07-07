package usecase

import (
	"github.com/enuesaa/walkhttp/pkg/config"
	"github.com/enuesaa/walkhttp/pkg/repository"
)

func ReadConfig(repos repository.Repos, path string) (config.Config, error) {
	configSrv := config.NewSrv(repos)
	return configSrv.Read(path)
}

func WriteConfig(repos repository.Repos, path string, conf config.Config) error {
	configSrv := config.NewSrv(repos)
	return configSrv.Write(path, conf)
}
