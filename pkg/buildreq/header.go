package buildreq

import (
	"fmt"

	"github.com/enuesaa/walkin/pkg/invoke"
)

func (b *Buildreq) AskHeader() (invoke.Header, error) {
	headerName := ""
	suggestion := []string{"content-type", "accept"}
	if err := b.repos.Prompt.AskSuggest("Header Name",  "(To skip, click enter)", suggestion, &headerName); err != nil {
		return invoke.Header{}, err
	}
	if headerName == "" {
		return invoke.Header{}, nil
	}

	headerValue := ""
	notice := fmt.Sprintf(" (%s) ", headerName)
	if err := b.repos.Prompt.Ask("Header Value",  notice, &headerValue); err != nil {
		return invoke.Header{}, err
	}
	header := invoke.Header{
		Key: headerName,
		Value: headerValue,
	}

	return header, nil
}
