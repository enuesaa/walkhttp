package usecase

import (
	"github.com/enuesaa/walkhttp/internal/invoke"
	"github.com/enuesaa/walkhttp/internal/repository"
)

func LoadConfig(repos repository.Repos, path string) invoke.Workspace {
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

func WriteConfig(repos repository.Repos, path string, conf invoke.Workspace) error {
	configSrv := invoke.New(repos)
	return configSrv.Write(path, conf)
}
