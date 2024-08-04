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
			wspath := c.String("workspace")
			ws := usecase.ReadWorkspace(repos, wspath)

			return usecase.Prompt(repos, "DELETE", ws)
		},
	}

	return cmd
}
