package usecase

import (
	"fmt"

	"github.com/enuesaa/walkin/pkg/buildreq"
	"github.com/enuesaa/walkin/pkg/invoke"
	"github.com/enuesaa/walkin/pkg/repository"
)


func PromptReqConfirmOnly(repos repository.Repos, invocation *invoke.Invocation) error {
	builder := buildreq.New(repos, invocation)

	fmt.Printf("***\n")
	fmt.Printf("* %s\n", builder.Endpoint())
	fmt.Printf("*\n")
	fmt.Printf("* [Headers]\n")

	for _, header := range invocation.RequestHeaders {
		fmt.Printf("* %s: %s\n", header.Key, header.Value)
	}

	if invocation.Method == "POST" || invocation.Method == "PUT" {
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
