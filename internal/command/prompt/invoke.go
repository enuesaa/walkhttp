package prompt

import (
	"github.com/enuesaa/walkhttp/internal/invoke"
	"github.com/enuesaa/walkhttp/internal/repository"
)

func Invoke(repos repository.Repos, entry *invoke.Entry) error {
	invokeSrv := invoke.New(repos)

	return invokeSrv.Invoke(entry)
}
