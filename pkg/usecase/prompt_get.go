package usecase

import (
	"fmt"

	"github.com/enuesaa/walkin/pkg/buildreq"
	"github.com/enuesaa/walkin/pkg/invoke"
	"github.com/enuesaa/walkin/pkg/repository"
)

func PromptGet(repos repository.Repos, url string) (invoke.Invocation, error) {
	builder := buildreq.New(repos)
	builder.Invocation.Method = "GET"
	builder.Invocation.Url = url

	if builder.IsUrlEmpty() {
		url, err := builder.AskUrl()
		if err != nil {
			return builder.Invocation, err
		}
		builder.Invocation.Url = url
	}
	fmt.Printf("***\n")
	fmt.Printf("* GET %s\n", url)
	fmt.Printf("*\n")
	fmt.Printf("* [Headers]\n")

	for {
		header, err := builder.AskHeader()
		if err != nil {
			return builder.Invocation, err
		} 
		if header.Key == "" {
			break
		}
		builder.Invocation.RequestHeaders = append(builder.Invocation.RequestHeaders, header)
		fmt.Printf("* %s: %s\n", header.Key, header.Value)
	}
	fmt.Printf("***\n")

	return builder.Invocation, nil
}
