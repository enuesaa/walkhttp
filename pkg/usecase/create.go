package usecase

import (
	"github.com/enuesaa/walkin/pkg/invoke"
	"github.com/enuesaa/walkin/pkg/repository"
)

func Create(repos repository.Repos, method string, url string) invoke.Invocation {
	invokeSrv := invoke.NewInvokeSrv(repos)
	return invokeSrv.Create(method, url)
}
