package usecase

import (
	"github.com/enuesaa/walkhttp/pkg/schema"
	"github.com/enuesaa/walkhttp/pkg/repository"
)

func LoadConfig(repos repository.Repos, path string) schema.Config {
	if path == "" {
		return schema.NewConfig()
	}
	configSrv := schema.NewConfigSrv(repos)
	conf, err := configSrv.Read(path)
	if err != nil {
		return schema.NewConfig()
	}
	return conf
}

func WriteConfig(repos repository.Repos, path string, conf schema.Config) error {
	configSrv := schema.NewConfigSrv(repos)
	return configSrv.Write(path, conf)
}
