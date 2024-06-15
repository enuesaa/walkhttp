package usecase

import (
	"github.com/enuesaa/walkin/pkg/repository"
)

func Prompt(repos repository.Repos) error {
	methods := []string{"GET", "POST", "PUT", "DELETE"}
	selected := ""
	if err := repos.Prompt.Select(methods, &selected); err != nil {
		return err
	}

	invocation := Create(repos, selected, "")
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

	return nil
}
