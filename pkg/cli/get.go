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
					Inline(true).
					WithTheme(huh.ThemeDracula())
				if err := urlPrompt.Run(); err != nil {
					return err
				}
			}
			fmt.Printf("GET %s\n", url)
			headers := map[string]string{}

			for {
				addHeader := false
				choicePrompt :=  huh.NewForm(huh.NewGroup(
					huh.NewConfirm().
						Title("Would you like to add Request Header ?").
						Affirmative("Add").
						Negative("Skip").
						Value(&addHeader).
						Inline(true),
					)).WithTheme(huh.ThemeDracula())
				if err := choicePrompt.Run(); err != nil {
					return err
				}
				if !addHeader {
					break
				}
	
				headerName := ""				
				headerNamePrompt := huh.NewForm(huh.NewGroup(
					huh.NewInput().
						Title(" Header Name ").
						Suggestions([]string{"content-type", "accept"}).
						Value(&headerName).
						Inline(true),
				)).WithKeyMap(keymap).WithTheme(huh.ThemeDracula())

				if err := headerNamePrompt.Run(); err != nil {
					return err
				}
				headerValue := ""
				headerValuePrompt := huh.NewForm(huh.NewGroup(
					huh.NewInput().
						Title(fmt.Sprintf(" Header Value [%s] ", headerName)).
						Value(&headerValue).
						Inline(true),
				)).WithKeyMap(keymap).WithTheme(huh.ThemeDracula())

				if err := headerValuePrompt.Run(); err != nil {
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
