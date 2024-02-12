package repository

import (
	"github.com/charmbracelet/huh"
)

type PromptInterface interface {
	Ask(message string, value *string) error
	Select(message string, choices []string, value *string) error
}
type Prompt struct{}

func (prompt *Prompt) Ask(message string, value *string) error {
	p := huh.NewInput().Title(message).Value(value)
	return p.Run()
}

func (promp *Prompt) Select(message string, choices []string, value *string) error {
	options := make([]huh.Option[string], 0)
	for _, choice := range choices {
		options = append(options, huh.NewOption[string](choice, choice))
	}

	p := huh.NewSelect[string]().Title(message).Options(options...).Value(value)

	return p.Run()
}
