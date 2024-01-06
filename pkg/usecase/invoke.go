package usecase

import (
	"fmt"

	"github.com/enuesaa/walkin/pkg/endpoint"
	"github.com/enuesaa/walkin/pkg/repository"
)

func Invoke(repos repository.Repos, name string) error {
	endpointSrv := endpoint.New(repos)
	configfile, err := endpointSrv.ReadConfig()
	if err != nil {
		return fmt.Errorf("failed to read config file")
	}

	endpointdef := endpoint.Endpoint{
		Name: "",
	}
	for _, def := range configfile.Endpoints {
		if def.Name == name {
			endpointdef = def
			break
		}
	}
	if endpointdef.Name == "" {
		return fmt.Errorf("failed to find endpoint with name %s", name)
	}
	fmt.Printf("%+v", endpointdef)

	return nil
}