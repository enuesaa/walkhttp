package repository

import (
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/huh"
)

type PromptInterface interface {
	Confirm(message string, value *bool) error
	Ask(message string, notice string, value *string) error
	AskSuggest(message string, notice string, suggestion []string, value *string) error
	Text(message string, notice string, value *string) error
}
type Prompt struct{}

func (prompt *Prompt) Confirm(message string, value *bool) error {
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().
				Title(message).
				Value(value),
		),
	)
	form.WithKeyMap(prompt.keymap())
	form.WithTheme(huh.ThemeBase())

	return form.Run()
}

func (prompt *Prompt) Ask(message string, notice string, value *string) error {
	description := " "
	if notice != "" {
		description = fmt.Sprintf(" %s ", notice)
	}
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewNote(),
			huh.NewInput().
				Title(message).
				Description(description).
				Value(value).
				Inline(true),
		),
	)
	form.WithKeyMap(prompt.keymap())
	form.WithTheme(huh.ThemeBase())

	return form.Run()
}

func (prompt *Prompt) AskSuggest(message string, notice string, suggestion []string, value *string) error {
	description := " "
	if notice != "" {
		description = fmt.Sprintf(" %s ", notice)
	}
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewNote(),
			huh.NewInput().
				Title(message).
				Description(description).
				Suggestions(suggestion).
				Value(value).
				Inline(true),
		),
	)
	form.WithKeyMap(prompt.keymap())
	form.WithTheme(huh.ThemeBase())

	return form.Run()
}

func (prompt *Prompt) Text(message string, notice string, value *string) error {
	description := " "
	if notice != "" {
		description = fmt.Sprintf(" %s ", notice)
	}
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewNote(),
			huh.NewText().
				Title(message).
				Description(description).
				Value(value),
		),
	)
	form.WithKeyMap(prompt.keymap())
	form.WithTheme(huh.ThemeBase())

	return form.Run()
}

func (prompt *Prompt) keymap() *huh.KeyMap {
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
	return keymap
}
