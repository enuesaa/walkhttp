package usecase

import (
	"fmt"

	"github.com/enuesaa/walkin/pkg/invoke"
	"github.com/enuesaa/walkin/pkg/repository"
)

func PromptReq(repos repository.Repos, invocation *invoke.Invocation) error {
	fmt.Printf("***\n")
	if invocation.Url == "" {
		invocation.Url = "https://"
		if err := repos.Prompt.Ask("Url", "", &invocation.Url); err != nil {
			return err
		}
	}
	fmt.Printf("* %s %s\n", invocation.Method, invocation.Url)
	fmt.Printf("*\n")
	fmt.Printf("* [Headers]\n")

	for {
		header := invoke.Header{}
		suggestion := []string{"content-type", "accept"}
		if err := repos.Prompt.AskSuggest("Header Name",  "(To skip, click enter)", suggestion, &header.Key); err != nil {
			return err
		}
		if header.Key == "" {
			break
		}

		if err := repos.Prompt.Ask("Header Value",  fmt.Sprintf(" (%s) ", header.Key), &header.Value); err != nil {
			return err
		}
		invocation.RequestHeaders = append(invocation.RequestHeaders, header)
		fmt.Printf("* %s: %s\n", header.Key, header.Value)
	}

	if invocation.Method == "POST" || invocation.Method == "PUT" {
		if err := repos.Prompt.Text("Body", "", &invocation.RequestBody); err != nil {
			return err
		}
		fmt.Printf("*\n")
		fmt.Printf("* [Body]\n")
		fmt.Printf("* %s\n", invocation.RequestBody)
	}
	fmt.Printf("***\n")

	confirm := true
	if err := repos.Prompt.Confirm("Do you confirm?", &confirm); err != nil {
		return err
	}
	if !confirm {
		return fmt.Errorf("unconfirmed")
	}

	return nil
}
