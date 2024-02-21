package usecase

import (
	"fmt"

	"github.com/enuesaa/walkin/pkg/buildreq"
	"github.com/enuesaa/walkin/pkg/invoke"
	"github.com/enuesaa/walkin/pkg/repository"
)


func PromptReq(repos repository.Repos, invocation *invoke.Invocation) error {
	builder := buildreq.New(repos, invocation)

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
		// TODO: remove builder
		if err := builder.AskHeader(); err != nil {
			if err == buildreq.SKIP_HEADER {
				break
			}
			return err
		}
		lastHeader := builder.GetLastHeader()
		fmt.Printf("* %s: %s\n", lastHeader.Key, lastHeader.Value)
	}
	if invocation.Method == "POST" || invocation.Method == "PUT" {
		fmt.Printf("*\n")

		body := ""
		if err := repos.Prompt.Text("Body", "", &body); err != nil {
			return err
		}
		invocation.RequestBody = []byte(body)

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
