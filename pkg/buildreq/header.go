package buildreq

import (
	"errors"
	"fmt"

	"github.com/enuesaa/walkin/pkg/invoke"
)

// almost same purpose as io.EOF
var SKIP_HEADER = errors.New("SKIP_HEADER")

func (b *Buildreq) AskHeader() error {
	headerName := ""
	suggestion := []string{"content-type", "accept"}
	if err := b.repos.Prompt.AskSuggest("Header Name",  "(To skip, click enter)", suggestion, &headerName); err != nil {
		return err
	}
	if headerName == "" {
		return SKIP_HEADER
	}

	headerValue := ""
	notice := fmt.Sprintf(" (%s) ", headerName)
	if err := b.repos.Prompt.Ask("Header Value",  notice, &headerValue); err != nil {
		return err
	}
	header := invoke.Header{
		Key: headerName,
		Value: headerValue,
	}
	b.Invocation.RequestHeaders = append(b.Invocation.RequestHeaders, header)

	return nil
}

func (b *Buildreq) GetLastHeader() invoke.Header {
	i := len(b.Invocation.RequestHeaders)
	if i == 0 {
		return invoke.Header{}
	}
	return b.Invocation.RequestHeaders[i - 1]
}
