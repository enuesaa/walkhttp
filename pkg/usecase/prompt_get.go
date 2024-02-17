package usecase

import (
	"fmt"

	"github.com/enuesaa/walkin/pkg/invoke"
	"github.com/enuesaa/walkin/pkg/repository"
)

func PromptGet(repos repository.Repos, url string) (invoke.Invocation, error) {
	invocation := invoke.Invocation {
		Method: "GET",
		Url: url,
		RequestHeaders: make([]invoke.Header, 0),
	}

	if url == "" {
		url = "https://"
		if err := repos.Prompt.Ask("Url", "", &url); err != nil {
			return invocation, err
		}
	}
	fmt.Printf("GET %s\n", url)

	for {
		headerName := ""
		suggestion := []string{"content-type", "accept"}
		if err := repos.Prompt.AskSuggest("Header Name",  "(To skip, click enter)", suggestion, &headerName); err != nil {
			return invocation, err
		}
		if headerName == "" {
			break
		}

		headerValue := ""
		notice := fmt.Sprintf(" (%s) ", headerName)
		if err := repos.Prompt.Ask("Header Value",  notice, &headerValue); err != nil {
			return invocation, err
		}
		invocation.RequestHeaders = append(invocation.RequestHeaders, invoke.Header{
			Key: headerName,
			Value: headerValue,
		})
		fmt.Printf("%s: %s\n", headerName, headerValue)
	}

	return invocation, nil
}
