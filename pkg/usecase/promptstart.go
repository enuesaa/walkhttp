package usecase

import (
	"github.com/enuesaa/walkin/pkg/repository"
)

func PromptStart(repos repository.Repos) error {
	methods := []string{"GET", "POST", "PUT", "DELETE"}
	selected := ""
	if err := repos.Prompt.Select(methods, &selected); err != nil {
		return err
	}

	invocation := Create(repos, selected, "")
	if err := PromptReq(repos, &invocation); err != nil {
		return err
	}

	return PromptStart(repos)
}
