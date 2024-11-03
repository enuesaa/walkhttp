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
		Action: func(c *cli.Context) error {
			return usecase.Prompt(repos, "GET")
		},
	}

	return cmd
}
