package usecase

import (
	"github.com/enuesaa/walkhttp/pkg/invoke"
	"github.com/enuesaa/walkhttp/pkg/repository"
)

func LoadConfig(repos repository.Repos, path string) invoke.Config {
	if path == "" {
		return invoke.NewConfig()
	}
	configSrv := invoke.New(repos)
	conf, err := configSrv.Read(path)
	if err != nil {
		return invoke.NewConfig()
	}
	return conf
}

func WriteConfig(repos repository.Repos, path string, conf invoke.Config) error {
	configSrv := invoke.New(repos)
	return configSrv.Write(path, conf)
}
