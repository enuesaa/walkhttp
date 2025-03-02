package prompt

import (
	"github.com/enuesaa/walkhttp/internal/invoke"
	"github.com/enuesaa/walkhttp/internal/repository"
)

func Prompt(repos repository.Repos) error {
	invokeSrv := invoke.New(repos)

	for {
		entry := invokeSrv.NewEntry("GET", "")
		repos.Log.Printf("\n")
	
		if err := BuildReq(repos, &entry); err != nil {
			return err
		}
		confirm := true
		if err := repos.Prompt.Confirm("Do you confirm?", &confirm); err != nil {
			return err
		}
		if !confirm {
			repos.Log.Printf("│ Canceled! \n")
			continue
		}
		if err := Invoke(repos, &entry); err != nil {
			return err
		}
	
		repos.Log.Printf("│ Status: %d\n", entry.Response.Status)
	
		if err := invokeSrv.Save(entry); err != nil {
			return err
		}
	}
}
