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
				url = "https://"
				if err := repos.Prompt.Ask("Url", "", &url); err != nil {
					return err
				}
			}
			fmt.Printf("GET %s\n", url)
			headers := map[string]string{}

			for {
				headerName := ""
				suggestion := []string{"content-type", "accept"}
				if err := repos.Prompt.AskSuggest("Header Name",  "(To skip, click enter)", suggestion, &headerName); err != nil {
					return err
				}
				if headerName == "" {
					break
				}

				headerValue := ""
				notice := fmt.Sprintf(" (%s) ", headerName)
				if err := repos.Prompt.Ask("Header Value",  notice, &headerValue); err != nil {
					return err
				}
				fmt.Printf("%s: %s\n", headerName, headerValue)
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
