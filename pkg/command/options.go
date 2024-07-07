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
				Name:     "collection",
				Aliases: []string{"c"},
				Usage:    "walkhttp collection file path",
				Value:    "walkhttp.json",
			},
		},
		Action: func(c *cli.Context) error {
			return usecase.Prompt(repos, "OPTIONS")
		},
	}

	return app
}
