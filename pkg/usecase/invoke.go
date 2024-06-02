package usecase

import (
	"github.com/enuesaa/walkin/pkg/invoke"
	"github.com/enuesaa/walkin/pkg/repository"
)

func Invoke(repos repository.Repos, invocation *invoke.Invocation) error {
	invokeSrv := invoke.NewInvokeSrv(repos)
	if err := invokeSrv.Invoke(invocation); err != nil {
		return err
	}
	repos.Log.Printf("status: %d\n", invocation.Status)
	return nil
}
