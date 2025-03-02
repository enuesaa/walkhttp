package command

import (
	"github.com/enuesaa/walkhttp/internal/invoke"
	"github.com/enuesaa/walkhttp/internal/repository"
)

func runPrompt(repos repository.Repos) error {
	invokeSrv := invoke.New(repos)

	for {
		entry := invokeSrv.NewEntry("GET", "")

		if err := buildReq(repos, &entry); err != nil {
			return err
		}
		if err := invokeSrv.Invoke(&entry); err != nil {
			return err
		}
	
		repos.Log.Printf("│ Status: %d\n", entry.Response.Status)
		repos.Log.Printf("\n")

		if err := invokeSrv.Save(entry); err != nil {
			return err
		}
	}
}

