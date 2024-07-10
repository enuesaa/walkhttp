package usecase

import (
	"github.com/enuesaa/walkhttp/pkg/invoke"
	"github.com/enuesaa/walkhttp/pkg/repository"
)

func Prompt(repos repository.Repos, method string, conf invoke.Config) error {
	invokeSrv := invoke.NewInvokeSrv(repos)
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

	if err := repos.DB.Save(invocation.ID, invocation); err != nil {
		return err
	}

	// config file modified
	// do you save?

	// リクエストが終わったら prompt でクエリできる
	// headers
	// headers.Accept
	// body

	return nil
}
