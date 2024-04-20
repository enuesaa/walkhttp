package usecase

import (
	"fmt"

	"github.com/enuesaa/walkin/pkg/invoke"
	"github.com/enuesaa/walkin/pkg/repository"
)

func PromptReqConfirmOnly(repos repository.Repos, invocation *invoke.Invocation) error {
	repos.Log.Printf("***\n")
	repos.Log.Printf("* %s %s\n", invocation.Method, invocation.Url)
	repos.Log.Printf("*\n")
	repos.Log.Printf("* [Headers]\n")

	for _, header := range invocation.RequestHeaders {
		repos.Log.Printf("* %s: %s\n", header.Key, header.Value)
	}

	if invocation.Method == "POST" || invocation.Method == "PUT" {
		repos.Log.Printf("*\n")
		repos.Log.Printf("* [Body]\n")
		repos.Log.Printf("* %s\n", invocation.RequestBody)
	}
	repos.Log.Printf("***\n")

	confirm := true
	if err := repos.Prompt.Confirm("Do you confirm?", &confirm); err != nil {
		return err
	}
	if !confirm {
		return fmt.Errorf("unconfirmed")
	}

	return nil
}
