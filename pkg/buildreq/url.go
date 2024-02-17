package buildreq

import "fmt"

func (b *Buildreq) IsUrlEmpty() bool {
	return b.Invocation.Url == ""
}

func (b *Buildreq) AskUrl() error {
	url := "https://"
	if err := b.repos.Prompt.Ask("Url", "", &url); err != nil {
		return err
	}
	b.Invocation.Url = url

	return nil
}

func (b *Buildreq) Endpoint() string {
	return fmt.Sprintf("%s %s", b.Invocation.Method, b.Invocation.Url)
}
