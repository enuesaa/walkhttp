package command

import (
	"github.com/enuesaa/walkhttp/pkg/repository"
	"github.com/enuesaa/walkhttp/pkg/usecase"
	"github.com/urfave/cli/v2"
)

func NewOptionsCommand(repos repository.Repos) *cli.Command {
	cmd := &cli.Command{
		Name:  "options",
		Usage: "make http OPTIONS request",
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

			return usecase.Prompt(repos, "OPTIONS", conf)
		},
	}

	return cmd
}
