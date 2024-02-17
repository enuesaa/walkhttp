package buildreq

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
