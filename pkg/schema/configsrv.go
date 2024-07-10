package schema

import (
	"encoding/json"

	"github.com/enuesaa/walkhttp/pkg/repository"
)

func NewConfigSrv(repos repository.Repos) ConfigSrv {
	return ConfigSrv{repos: repos}
}

type ConfigSrv struct {
	repos repository.Repos
}

func (srv *ConfigSrv) Read(path string) (Config, error) {
	fbytes, err := srv.repos.Fs.Read(path)
	if err != nil {
		return Config{}, err
	}
	var config Config
	if err := json.Unmarshal(fbytes, &config); err != nil {
		return Config{}, err
	}
	return config, nil
}

func (srv *ConfigSrv) Write(path string, config Config) error {
	fbytes, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}
	if err := srv.repos.Fs.Create(path, fbytes); err != nil {
		return  err
	}
	return nil
}
