package command

import (
	"github.com/enuesaa/walkhttp/internal/repository"
	"github.com/enuesaa/walkhttp/internal/usecase"
	"github.com/urfave/cli/v2"
)

func NewDeleteCommand(repos repository.Repos) *cli.Command {
	cmd := &cli.Command{
		Name:  "delete",
		Usage: "make http DELETE request",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "workspace",
				Aliases: []string{"w"},
				Usage:   "workspace file path",
				Value:   "walkhttp.json",
			},
		},
		Action: func(c *cli.Context) error {
			repos.WorkspacePath = c.String("workspace")

			return usecase.Prompt(repos, "DELETE")
		},
	}

	return cmd
}
