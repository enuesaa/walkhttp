package command

import (
	"github.com/enuesaa/walkhttp/internal/repository"
	"github.com/enuesaa/walkhttp/internal/usecase"
	"github.com/urfave/cli/v2"
)

func NewPostCommand(repos repository.Repos) *cli.Command {
	cmd := &cli.Command{
		Name:  "post",
		Usage: "make http POST request",
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
			repos.Ws.Use(wspath)

			return usecase.Prompt(repos, "POST")
		},
	}

	return cmd
}
