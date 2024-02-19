package buildreq

import (
	"github.com/enuesaa/walkin/pkg/invoke"
	"github.com/enuesaa/walkin/pkg/repository"
)

func New(repos repository.Repos, invocation *invoke.Invocation) Buildreq {
	return Buildreq{
		repos: repos,
		Invocation: *invocation,
	}
}

type Buildreq struct {
	repos repository.Repos
	Invocation invoke.Invocation
}