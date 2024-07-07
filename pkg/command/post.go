package command

import (
	"github.com/enuesaa/walkhttp/pkg/repository"
	"github.com/enuesaa/walkhttp/pkg/usecase"
	"github.com/urfave/cli/v2"
)

func NewPostCommand(repos repository.Repos) *cli.Command {
	app := &cli.Command{
		Name:  "post",
		Usage: "make http POST request",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "config",
				Aliases: []string{"c"},
				Usage:    "config file path",
				Value:    "walkhttp.json",
			},
		},
		Action: func(c *cli.Context) error {
			return usecase.Prompt(repos, "POST")
		},
	}

	return app
}