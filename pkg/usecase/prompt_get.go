package usecase

import (
	"fmt"

	"github.com/enuesaa/walkin/pkg/buildreq"
	"github.com/enuesaa/walkin/pkg/invoke"
	"github.com/enuesaa/walkin/pkg/repository"
)

func PromptGet(repos repository.Repos, url string) (invoke.Invocation, error) {
	builder := buildreq.New(repos, "GET", url)

	fmt.Printf("***\n")
	if builder.IsUrlEmpty() {
		if err := builder.AskUrl(); err != nil {
			return builder.Invocation, err
		}
	}
	fmt.Printf("* GET %s\n", builder.Invocation.Url)
	fmt.Printf("*\n")
	fmt.Printf("* [Headers]\n")

	for {
		if err := builder.AskHeader(); err != nil {
			if err == buildreq.SKIP_HEADER {
				break
			}
			return builder.Invocation, err
		}
		lastHeader := builder.GetLastHeader()
		fmt.Printf("* %s: %s\n", lastHeader.Key, lastHeader.Value)
	}
	fmt.Printf("***\n")

	return builder.Invocation, nil
}
