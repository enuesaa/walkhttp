package command

import (
	"github.com/enuesaa/walkhttp/internal/repository"
	"github.com/enuesaa/walkhttp/internal/usecase"
	"github.com/urfave/cli/v2"
)

func NewGetCommand(repos repository.Repos) *cli.Command {
	cmd := &cli.Command{
		Name:  "get",
		Usage: "make http GET request",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "workspace",
				Aliases: []string{"w"},
				Usage:   "workspace file path",
				Value:   "walkhttp.json",
			},
		},
		Action: func(c *cli.Context) error {
			configpath := c.String("workspace")
			conf := usecase.LoadConfig(repos, configpath)

			return usecase.Prompt(repos, "GET", conf)
		},
	}

	return cmd
}
