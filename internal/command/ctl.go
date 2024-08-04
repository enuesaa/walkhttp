package command

import (
	"fmt"
	"time"

	"github.com/enuesaa/walkhttp/internal/repository"
	"github.com/enuesaa/walkhttp/internal/serve"
	"github.com/enuesaa/walkhttp/internal/usecase"
	"github.com/pkg/browser"
	"github.com/urfave/cli/v2"
)

func NewCtlCommand(repos repository.Repos) *cli.Command {
	cmd := &cli.Command{
		Name:  "ctl",
		Usage: "serve web console",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:  "port",
				Usage: "Serve port",
				Value: 3000,
			},
			&cli.StringFlag{
				Name:    "workspace",
				Aliases: []string{"w"},
				Usage:   "workspace file path",
				Value:   "walkhttp.json",
			},
		},
		Action: func(c *cli.Context) error {
			port := c.Int("port")
			wspath := c.String("workspace")
			repos.Ws.Use(wspath)

			usecase.PrintBanner(repos)

			go func() {
				time.Sleep(1 * time.Second)
				url := fmt.Sprintf("http://localhost:%d", port)
				browser.OpenURL(url)
			}()

			return serve.Serve(repos, port)
		},
	}

	return cmd
}
