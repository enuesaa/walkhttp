package usecase

import (
	"github.com/enuesaa/walkin/pkg/endpoint"
	"github.com/enuesaa/walkin/pkg/repository"
)

func CreateConfigFileIfNotExist(repos repository.Repos) error {
	endpointSrv := endpoint.New(repos)
	if endpointSrv.IsConfigFileExist() {
		return nil
	}

	configfile := endpoint.ConfigFile{
		Endpoints: make(map[string]endpoint.Endpoint, 0),
	}
	return endpointSrv.WriteConfig(configfile)
}
