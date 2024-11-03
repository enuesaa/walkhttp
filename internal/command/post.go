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
		Action: func(c *cli.Context) error {
			return usecase.Prompt(repos, "POST")
		},
	}

	return cmd
}
