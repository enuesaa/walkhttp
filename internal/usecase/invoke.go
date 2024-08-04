package usecase

import (
	"github.com/enuesaa/walkhttp/internal/invoke"
	"github.com/enuesaa/walkhttp/internal/repository"
)

func Invoke(repos repository.Repos, invocation *invoke.Invocation) error {
	invokeSrv := invoke.New(repos)
	if err := invokeSrv.Invoke(invocation); err != nil {
		return err
	}
	return nil
}
