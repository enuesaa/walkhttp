package buildreq

import (
	"fmt"

	"github.com/enuesaa/walkin/pkg/invoke"
)

func (b *Buildreq) AskHeader() (bool, error) {
	headerName := ""
	suggestion := []string{"content-type", "accept"}
	if err := b.repos.Prompt.AskSuggest("Header Name",  "(To skip, click enter)", suggestion, &headerName); err != nil {
		return false, err
	}
	if headerName == "" {
		return false, nil
	}

	headerValue := ""
	notice := fmt.Sprintf(" (%s) ", headerName)
	if err := b.repos.Prompt.Ask("Header Value",  notice, &headerValue); err != nil {
		return false, err
	}
	header := invoke.Header{
		Key: headerName,
		Value: headerValue,
	}
	b.Invocation.RequestHeaders = append(b.Invocation.RequestHeaders, header)

	return false, nil
}

func (b *Buildreq) GetLastHeader() invoke.Header {
	i := len(b.Invocation.RequestHeaders)
	if i == 0 {
		return invoke.Header{}
	}
	return b.Invocation.RequestHeaders[i - 1]
}
