package config

import (
	"encoding/json"

	"github.com/enuesaa/walkin/pkg/repository"
)

func NewConfigSrv(repos repository.Repos) ConfigSrv {
	return ConfigSrv{
		repos: repos,
	}
}

type ConfigSrv struct {
	repos repository.Repos
}
func (srv *ConfigSrv) Write(config Config) error {
	fbytes, err := json.Marshal(config)
	if err != nil {
		return err
	}
	return srv.repos.Fs.Create("walkin.json", fbytes)
}

func (srv *ConfigSrv) IsExist() bool {
	return srv.repos.Fs.IsExist("walkin.json")
}

func (srv *ConfigSrv) Read() (Config, error) {
	fbytes, err := srv.repos.Fs.Read("walkin.json")
	if err != nil {
		return Config{}, err
	}
	var config Config
	if err := json.Unmarshal(fbytes, &config); err != nil {
		return Config{}, err
	}
	return config, nil
}

