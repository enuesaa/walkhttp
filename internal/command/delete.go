package command

import (
	"github.com/enuesaa/walkhttp/internal/repository"
	"github.com/enuesaa/walkhttp/internal/command/prompt"
	"github.com/urfave/cli/v2"
)

func NewDeleteCommand(repos repository.Repos) *cli.Command {
	cmd := &cli.Command{
		Name:  "delete",
		Usage: "make http DELETE request",
		Action: func(c *cli.Context) error {
			return prompt.Prompt(repos, "DELETE")
		},
	}

	return cmd
}
