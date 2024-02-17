package usecase

import (
	"fmt"

	"github.com/enuesaa/walkin/pkg/buildreq"
	"github.com/enuesaa/walkin/pkg/invoke"
	"github.com/enuesaa/walkin/pkg/repository"
)

func PromptGet(repos repository.Repos, url string) (invoke.Invocation, error) {
	builder := buildreq.New(repos, "GET", url)
	builder.Debug = true

	if builder.IsUrlEmpty() {
		if err := builder.AskUrl(); err != nil {
			return builder.Invocation, err
		}
	}
	fmt.Printf("***\n")
	fmt.Printf("* GET %s\n", builder.Invocation.Url)
	fmt.Printf("*\n")
	fmt.Printf("* [Headers]\n")

	for {
		skipped, err := builder.AskHeader()
		if err != nil {
			return builder.Invocation, err
		}
		if skipped {
			break
		}
		lastHeader := builder.GetLastHeader()
		fmt.Printf("* %s: %s\n", lastHeader.Key, lastHeader.Value)
	}
	fmt.Printf("***\n")

	return builder.Invocation, nil
}
