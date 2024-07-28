package command

import (
	"fmt"
	"time"

	"github.com/enuesaa/walkhttp/pkg/repository"
	"github.com/enuesaa/walkhttp/pkg/serve"
	"github.com/enuesaa/walkhttp/pkg/usecase"
	"github.com/pkg/browser"
	"github.com/urfave/cli/v2"
)

func NewCtlCommand(repos repository.Repos) *cli.Command {
	cmd := &cli.Command{
		Name:  "ctl",
		Usage: "Serve web console",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:  "port",
				Usage: "Serve port",
				Value: 3000,
			},
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Usage:   "config file path",
				Value:   "walkhttp.json",
			},
		},
		Action: func(c *cli.Context) error {
			port := c.Int("port")
			configpath := c.String("config")
			conf := usecase.LoadConfig(repos, configpath)

			usecase.PrintBanner(repos)

			go func() {
				time.Sleep(1 * time.Second)
				url := fmt.Sprintf("http://localhost:%d", port)
				browser.OpenURL(url)
			}()

			return serve.Serve(repos, conf.BaseUrl, port)
		},
	}

	return cmd
}
