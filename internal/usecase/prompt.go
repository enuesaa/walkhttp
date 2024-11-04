package usecase

import (
	"github.com/enuesaa/walkhttp/internal/invoke"
	"github.com/enuesaa/walkhttp/internal/repository"
)

func Prompt(repos repository.Repos, method string) error {
	invokeSrv := invoke.New(repos)
	entry := invokeSrv.NewEntry(method)

	repos.Log.Printf("***\n")
	defer repos.Log.Printf("***\n")
	if err := PromptReq(repos, &entry); err != nil {
		return err
	}

	if err := Invoke(repos, &entry); err != nil {
		return err
	}
	repos.Log.Printf("* Status: %d\n", entry.Response.Status)

	return invokeSrv.Write(entry)
}
