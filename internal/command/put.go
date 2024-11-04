package command

import (
	"github.com/enuesaa/walkhttp/internal/repository"
	"github.com/enuesaa/walkhttp/internal/command/prompt"
	"github.com/urfave/cli/v2"
)

func NewPutCommand(repos repository.Repos) *cli.Command {
	cmd := &cli.Command{
		Name:  "put",
		Usage: "make http PUT request",
		Action: func(c *cli.Context) error {
			return prompt.Prompt(repos, "PUT")
		},
	}

	return cmd
}
