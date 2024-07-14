package usecase

import (
	"github.com/enuesaa/walkhttp/pkg/invoke"
	"github.com/enuesaa/walkhttp/pkg/repository"
	"github.com/enuesaa/walkhttp/pkg/serve/schema"
)

func Invoke(repos repository.Repos, invocation *schema.Invocation) error {
	invokeSrv := invoke.New(repos)
	if err := invokeSrv.Invoke(invocation); err != nil {
		return err
	}
	return nil
}
