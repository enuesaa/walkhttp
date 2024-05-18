package usecase

import (
	"fmt"

	"github.com/enuesaa/walkin/pkg/invoke"
	"github.com/enuesaa/walkin/pkg/repository"
)

func PromptReq(repos repository.Repos, invocation *invoke.Invocation) error {
	repos.Log.Printf("***\n")
	if err := repos.Prompt.Ask("Url", "", &invocation.URL); err != nil {
		return err
	}
	repos.Log.Printf("* %s %s\n", invocation.Method, invocation.URL)
	repos.Log.Printf("*\n")
	repos.Log.Printf("* [Headers]\n")

	for {
		header := invoke.Header{}
		suggestion := []string{"content-type", "accept"}
		if err := repos.Prompt.AskSuggest("Header Name", "(To skip, click enter)", suggestion, &header.Name); err != nil {
			return err
		}
		if header.Name == "" {
			break
		}

		if err := repos.Prompt.Ask("Header Value", fmt.Sprintf(" (%s) ", header.Name), &header.Value); err != nil {
			return err
		}
		invocation.RequestHeaders = append(invocation.RequestHeaders, &header)
		repos.Log.Printf("* %s: %s\n", header.Name, header.Value)
	}

	if invocation.Method == "POST" || invocation.Method == "PUT" {
		if err := repos.Prompt.Text("Body", "", &invocation.RequestBody); err != nil {
			return err
		}
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
