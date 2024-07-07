package cli

import (
	"github.com/enuesaa/walkhttp/pkg/repository"
	"github.com/urfave/cli/v2"
)

func NewPostCommand(repos repository.Repos) *cli.Command {
	app := &cli.Command{
		Name:  "post",
		Usage: "make http POST request",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "collection",
				Aliases: []string{"c"},
				Usage:    "walkhttp collection file path",
				Value:    "walkhttp.json",
			},
		},
		Action: func(c *cli.Context) error {
			return nil
		},
	}

	return app
}
