package usecase

import (
	"github.com/enuesaa/walkhttp/pkg/invoke"
	"github.com/enuesaa/walkhttp/pkg/serve/schema"
	"github.com/enuesaa/walkhttp/pkg/repository"
)

func Invoke(repos repository.Repos, invocation *schema.Invocation) error {
	invokeSrv := invoke.NewInvokeSrv(repos)
	if err := invokeSrv.Invoke(invocation); err != nil {
		return err
	}
	return nil
}
