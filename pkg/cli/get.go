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
				urlPrompt := huh.NewInput().Title("Url").Value(&url)
				if err := urlPrompt.Run(); err != nil {
					return err
				}
			}

			headers := map[string]string{}

			for {
				task := ""
				taskPrompt := huh.NewSelect[string]().
					Title("Choose your burger").
					Options(
						huh.NewOption[string]("Call", "call"),
						huh.NewOption[string]("Header", "header"),
					).
					Value(&task)
				if err := taskPrompt.Run(); err != nil {
					return err
				}

				if task == "call" {
					break
				}
				if task == "header" {
					headerName := ""
					headerValue := ""
					headerPrompt := huh.NewForm(
						huh.NewGroup(
							huh.NewInput().
								Title("Header name").
								Suggestions([]string{"content-type", "accept"}).
								Value(&headerName),
							huh.NewInput().
								Title("Header value").
								Value(&headerValue),
						),
					)
					if err := headerPrompt.WithKeyMap(keymap).Run(); err != nil {
						return err
					}
					headers[headerName] = headerValue
				}
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
