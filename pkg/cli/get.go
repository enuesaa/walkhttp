package cli

import (
	"fmt"

	"github.com/enuesaa/walkin/pkg/invoke"
	"github.com/enuesaa/walkin/pkg/repository"
	"github.com/spf13/cobra"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/huh"
)

func CreateGetCmd(repos repository.Repos) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "get request",
		RunE: func(cmd *cobra.Command, args []string) error {
			url, _ := cmd.Flags().GetString("url")

			keymap := huh.NewDefaultKeyMap()
			keymap.Input.AcceptSuggestion = key.NewBinding(
				key.WithKeys("tab"),
			)
			keymap.Input.Prev = key.NewBinding(
				key.WithKeys("up"),
				key.WithHelp("up", "back"),
			)
			keymap.Input.Next = key.NewBinding(
				key.WithKeys("enter"),
			)
			keymap.Input.Submit = key.NewBinding(
				key.WithKeys("enter"),
			)

			if url == "" {
				url = "https://" // default value
				urlPrompt := huh.NewInput().
					Title("Url ").
					Value(&url).
					Inline(true)
				if err := urlPrompt.Run(); err != nil {
					return err
				}
			}
			fmt.Printf("GET %s\n", url)

			headers := map[string]string{}

			for {
				headerName := ""				
				headerNamePrompt := huh.NewForm(huh.NewGroup(
					huh.NewInput().
						Title("Header Name").
						Description(" (skip if empty) ").
						Suggestions([]string{"content-type", "accept"}).
						Value(&headerName).
						Inline(true),
				)).WithKeyMap(keymap)

				if err := headerNamePrompt.Run(); err != nil {
					return err
				}
				if headerName == "" {
					break
				}
				headerValue := ""
				headerValuePrompt := huh.NewInput().
					Title(fmt.Sprintf("Header Value %s", headerName)).
					Value(&headerValue).
					Inline(true).
					WithKeyMap(keymap)
				if err := headerValuePrompt.Run(); err != nil {
					return err
				}
				headers[headerName] = headerValue
			}

			invocation := invoke.Invocation {
				Method: "GET",
				Url: url,
			}
			if err := invoke.Invoke(&invocation); err != nil {
				return err
			}
			fmt.Printf("status: %d\n", invocation.Status)
			fmt.Printf("body: %s\n", string(invocation.ResponseBody))

			return nil
		},
	}
	cmd.Flags().String("url", "", "url")

	return cmd
}
