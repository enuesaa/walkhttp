package buildreq

func (b *Buildreq) AskBody() error {
	body := ""
	if err := b.repos.Prompt.Text("Body", "", &body); err != nil {
		return err
	}
	b.Invocation.RequestBody = []byte(body)

	return nil
}
