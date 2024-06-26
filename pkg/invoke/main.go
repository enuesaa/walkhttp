package invoke

import (
	"github.com/enuesaa/walkhttp/pkg/repository"
)

func NewInvokeSrv(repos repository.Repos) InvokeSrv {
	return InvokeSrv{
		repos: repos,
	}
}

type InvokeSrv struct {
	repos repository.Repos
}
