package invoke

import (
	"strings"

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

func (srv *InvokeSrv) BaseUrl() string {
	base := srv.repos.Env.WALKHTTP_URL_BASE

	if strings.HasSuffix(base, "/") {
		return strings.TrimSuffix(base, "/")
	}
	return base
}
