package usecase

import (
	"github.com/enuesaa/walkhttp/pkg/config"
	"github.com/enuesaa/walkhttp/pkg/repository"
)

func LoadConfig(repos repository.Repos, path string) config.Config {
	if path == "" {
		return config.New()
	}
	configSrv := config.NewSrv(repos)
	conf, err := configSrv.Read(path)
	if err != nil {
		return config.New()
	}
	return conf
}

func WriteConfig(repos repository.Repos, path string, conf config.Config) error {
	configSrv := config.NewSrv(repos)
	return configSrv.Write(path, conf)
}
