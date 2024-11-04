package command

import (
	"github.com/enuesaa/walkhttp/internal/repository"
	"github.com/enuesaa/walkhttp/internal/command/prompt"
	"github.com/urfave/cli/v2"
)

func NewOptionsCommand(repos repository.Repos) *cli.Command {
	cmd := &cli.Command{
		Name:  "options",
		Usage: "make http OPTIONS request",
		Action: func(c *cli.Context) error {
			return prompt.Prompt(repos, "OPTIONS")
		},
	}

	return cmd
}
