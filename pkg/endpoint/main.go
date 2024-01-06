package endpoint

import (
	"github.com/enuesaa/walkin/pkg/repository"
)

func New(repos repository.Repos) EndpointSrv {
	return EndpointSrv{
		repos: repos,
	}
}

type EndpointSrv struct {
	repos repository.Repos
}
