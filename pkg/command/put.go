package command

import (
	"github.com/enuesaa/walkhttp/pkg/repository"
	"github.com/enuesaa/walkhttp/pkg/usecase"
	"github.com/urfave/cli/v2"
)

func NewPutCommand(repos repository.Repos) *cli.Command {
	cmd := &cli.Command{
		Name:  "put",
		Usage: "make http PUT request",
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

			return usecase.Prompt(repos, "PUT", conf)
		},
	}

	return cmd
}
