package invoke

import (
	"github.com/enuesaa/walkhttp/internal/repository"
)

func New(repos repository.Repos) InvokeSrv {
	return InvokeSrv{
		repos: repos,
	}
}

type InvokeSrv struct {
	repos repository.Repos
}
