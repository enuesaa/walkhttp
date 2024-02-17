package buildreq

import (
	"github.com/enuesaa/walkin/pkg/invoke"
	"github.com/enuesaa/walkin/pkg/repository"
)

func New(repos repository.Repos, method string, url string) Buildreq {
	return Buildreq{
		repos: repos,
		Debug: false,
		Invocation: invoke.Invocation {
			Method: method,
			Url: url,
			RequestHeaders: make([]invoke.Header, 0),
		},
	}
}

type Buildreq struct {
	repos repository.Repos
	Debug bool
	Invocation invoke.Invocation
}