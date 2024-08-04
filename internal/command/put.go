package command

import (
	"github.com/enuesaa/walkhttp/internal/repository"
	"github.com/enuesaa/walkhttp/internal/usecase"
	"github.com/urfave/cli/v2"
)

func NewPutCommand(repos repository.Repos) *cli.Command {
	cmd := &cli.Command{
		Name:  "put",
		Usage: "make http PUT request",
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

			return usecase.Prompt(repos, "PUT", conf)
		},
	}

	return cmd
}
