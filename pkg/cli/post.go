package cli

import (
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/huh"
	"github.com/enuesaa/walkin/pkg/repository"
	"github.com/spf13/cobra"
)

func CreatePostCmd(repos repository.Repos) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "post",
		Short: "post request",
		RunE: func(cmd *cobra.Command, args []string) error {
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

			url := ""
			urlPrompt := huh.NewInput().Title("Url").Value(&url).WithKeyMap(keymap)
			if err := urlPrompt.Run(); err != nil {
				return err
			}

			for {
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
				headerPrompt.WithKeyMap(keymap)

				if err := headerPrompt.Run(); err != nil {
					return err
				}

				break
			}
			fmt.Printf("url: %s\n", url)

			return nil
		},
	}

	return cmd
}
