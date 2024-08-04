package usecase

import (
	"github.com/enuesaa/walkhttp/internal/invoke"
	"github.com/enuesaa/walkhttp/internal/repository"
	"github.com/enuesaa/walkhttp/internal/repository/workspace"
)

func Prompt(repos repository.Repos, method string, conf workspace.Workspace) error {
	invokeSrv := invoke.New(repos)
	invocation := invokeSrv.Create(method, conf.BaseUrl)
	repos.Log.Printf("***\n")
	if err := PromptReq(repos, &invocation); err != nil {
		repos.Log.Printf("***\n")
		return err
	}
	if err := Invoke(repos, &invocation); err != nil {
		repos.Log.Printf("***\n")
		return err
	}
	repos.Log.Printf("* Status: %d\n", invocation.Status)
	repos.Log.Printf("***\n")
	repos.Log.Printf("\n")

	return invokeSrv.Save(invocation)
}
