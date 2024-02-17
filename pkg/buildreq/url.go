package buildreq

func (b *Buildreq) IsUrlEmpty() bool {
	return b.Invocation.Url == ""
}

func (b *Buildreq) AskUrl() (string, error) {
	url := "https://"
	if err := b.repos.Prompt.Ask("Url", "", &url); err != nil {
		return "", err
	}
	return url, nil
}
