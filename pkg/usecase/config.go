package usecase

import (
	"github.com/enuesaa/walkhttp/pkg/config"
	"github.com/enuesaa/walkhttp/pkg/repository"
)

func LoadConfig(repos repository.Repos, path string) config.Config {
	if path == "" {
		return config.NewConfig()
	}
	configSrv := config.NewConfigSrv(repos)
	conf, err := configSrv.Read(path)
	if err != nil {
		return config.NewConfig()
	}
	return conf
}

func WriteConfig(repos repository.Repos, path string, conf config.Config) error {
	configSrv := config.NewConfigSrv(repos)
	return configSrv.Write(path, conf)
}
