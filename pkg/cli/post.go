package cli

import (
	"fmt"
	"log"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/huh"
	"github.com/enuesaa/walkin/pkg/repository"
	"github.com/spf13/cobra"
)

var (
    burger string
    toppings []string
    sauceLevel int
    name string
    instructions string
    discount bool
)

func CreatePostCmd(repos repository.Repos) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "post",
		Short: "post request",
		Run: func(cmd *cobra.Command, args []string) {
			form := huh.NewForm(
				huh.NewGroup(
					huh.NewSelect[string]().
						Title("Choose your burger").
						Options(
							huh.NewOption("Charmburger Classic", "classic"),
							huh.NewOption("Chickwich", "chickwich"),
						).
						Value(&burger),
					huh.NewMultiSelect[string]().
						Title("Toppings").
						Options(
							huh.NewOption("Lettuce", "lettuce").Selected(true),
							huh.NewOption("Jalapeños", "jalapeños"),
							huh.NewOption("Cheese", "cheese"),
						).
						Limit(2).
						Value(&toppings),
				),

				huh.NewGroup(
					huh.NewInput().
						Title("What's your name?").
						Suggestions([]string{"hey", "hello"}).
						Value(&name),
					huh.NewText().
						Title("Special Instructions").
						CharLimit(400).
						Value(&instructions),
					huh.NewConfirm().
						Title("Would you like 15% off?").
						Value(&discount),
				),
			)

			keymap := huh.NewDefaultKeyMap()
			keymap.Input.AcceptSuggestion = key.NewBinding(key.WithKeys("tab"), key.WithHelp("tab", "complete"))
			form.WithKeyMap(keymap)

			if err := form.Run(); err != nil {
				log.Fatalf("Error: %s", err.Error())
			}
			fmt.Printf("name: %s\n", name)
		},
	}

	return cmd
}
