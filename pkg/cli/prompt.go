package cli

import (
	"fmt"
	"strings"

	"github.com/enuesaa/walkhttp/pkg/repository"
	"github.com/spf13/cobra"

	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

// see https://github.com/charmbracelet/bubbletea/blob/master/examples/textinputs/main.go
// see https://github.com/charmbracelet/bubbletea/blob/master/examples/list-simple/main.go
func Prompt(repos repository.Repos) *cobra.Command {
	cmd := &cobra.Command{
		Use: "prompt",
		RunE: func(cmd *cobra.Command, args []string) error {
			selected := ""
		
			form := huh.NewSelect[string]()
			form.Options(
				huh.Option[string]{Key: "GET", Value: "GET"},
				huh.Option[string]{Key: "POST", Value: "POST"},
				huh.Option[string]{Key: "PUT", Value: "PUT"},
				huh.Option[string]{Key: "DELETE", Value: "DELETE"},
				huh.Option[string]{Key: "OPTIONS", Value: "OPTIONS"},
			)
			form.Value(&selected)
			form.Inline(true)
			if err := form.Run(); err != nil {
				return err
			}

			_, err := tea.NewProgram(initialModel()).Run()
			return err
		},
	}

	return cmd
}

var (
	focusedStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	blurredStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	noStyle       = lipgloss.NewStyle()
	focusedButton = focusedStyle.Render("[ Submit ]")
	blurredButton = fmt.Sprintf("[ %s ]", blurredStyle.Render("Submit"))
)

type model struct {
	focusIndex int
	inputs     []textinput.Model
}

func initialModel() model {
	m := model{
		inputs: make([]textinput.Model, 3),
	}

	var t textinput.Model
	for i := range m.inputs {
		t = textinput.New()
		t.Cursor.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
		t.Cursor.SetMode(cursor.CursorStatic)
		t.CharLimit = 32

		switch i {
		case 0:
			t.Placeholder = "Url"
			t.CharLimit = 64
		case 1:
			t.Placeholder = "Body"
		}
		m.inputs[i] = t
	}

	return m
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit
		case "tab", "shift+tab", "enter", "up", "down":
			s := msg.String()

			if s == "enter" && m.focusIndex == len(m.inputs) {
				return m, tea.Quit
			}

			if s == "up" || s == "shift+tab" {
				m.focusIndex--
			} else {
				m.focusIndex++
			}

			if m.focusIndex > len(m.inputs) {
				m.focusIndex = 0
			} else if m.focusIndex < 0 {
				m.focusIndex = len(m.inputs)
			}

			cmds := make([]tea.Cmd, len(m.inputs))
			for i := 0; i <= len(m.inputs)-1; i++ {
				if i == m.focusIndex {
					cmds[i] = m.inputs[i].Focus()
					m.inputs[i].PromptStyle = focusedStyle
					m.inputs[i].TextStyle = focusedStyle
				} else {
					m.inputs[i].Blur()
					m.inputs[i].PromptStyle = noStyle
					m.inputs[i].TextStyle = noStyle
				}
			}

			return m, tea.Batch(cmds...)
		}
	}
	cmd := m.updateInputs(msg)

	return m, cmd
}

func (m *model) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.inputs))

	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}

func (m model) View() string {
	var b strings.Builder

	for i := range m.inputs {
		b.WriteString(m.inputs[i].View())
		if i < len(m.inputs)-1 {
			b.WriteRune('\n')
		}
	}

	button := &blurredButton
	if m.focusIndex == len(m.inputs) {
		button = &focusedButton
	}
	fmt.Fprintf(&b, "\n\n%s\n\n", *button)

	return b.String()
}
