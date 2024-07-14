package command

import (
	"github.com/enuesaa/walkhttp/pkg/repository"
	"github.com/enuesaa/walkhttp/pkg/usecase"
	"github.com/urfave/cli/v2"
)

func NewOptionsCommand(repos repository.Repos) *cli.Command {
	app := &cli.Command{
		Name:  "options",
		Usage: "make http OPTIONS request",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Usage:   "config file path",
				Value:   "walkhttp.json",
			},
		},
		Action: func(c *cli.Context) error {
			configpath := c.String("config")
			conf := usecase.LoadConfig(repos, configpath)

			return usecase.Prompt(repos, "OPTIONS", conf)
		},
	}

	return app
}
