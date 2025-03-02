package command

import (
	"fmt"
	"log"

	"github.com/enuesaa/walkhttp/internal/repository"
	"github.com/enuesaa/walkhttp/internal/router"
	"github.com/urfave/cli/v2"
)

func New(repos repository.Repos) *cli.App {
	app := &cli.App{
		Name:    "walkhttp",
		Version: "0.0.11",
		Usage:   "A CLI tool to proxy HTTP requests for debugging.",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:  "port",
				Usage: "Port",
				Value: 3000,
			},
			&cli.StringFlag{
				Name:  "origin",
				Usage: "Origin URL",
				DefaultText: "https://example.com",
			},
			&cli.BoolFlag{
				Name:  "prompt",
				Usage: "Start prompt and call a HTTP request",
			},
		},
		Action: func(c *cli.Context) error {
			port := c.Int("port")
			origin := c.String("origin")
			prompt := c.Bool("prompt")
			if origin != "" {
				repos.Env.WALKHTTP_URL_BASE = origin
			}
			printBanner(repos, port)

			if prompt {
				go func () {
					prompter := Prompter{repos}
					if err := prompter.Run(); err != nil {
						log.Fatalf("Error: %s", err.Error())
					}
				}()
			}
			server := router.New(repos)
			address := fmt.Sprintf(":%d", port)

			return server.Start(address)
		},
	}

	// disable default
	app.OnUsageError = func(c *cli.Context, err error, isSubcommand bool) error {
		return err
	}
	app.HideHelpCommand = true
	app.Suggest = false
	cli.AppHelpTemplate = `{{.Usage}}

USAGE:
	{{.HelpName}} {{if .VisibleFlags}}[flags]{{end}}{{if .Commands}} command [command options]{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments...]{{end}}
	{{if len .Authors}}
AUTHOR:
	{{range .Authors}}{{ . }}{{end}}
	{{end}}{{if .Commands}}
COMMANDS:
{{range .Commands}}{{if not .HideHelp}}   {{join .Names ", "}}{{ "\t"}}{{.Usage}}{{ "\n" }}{{end}}{{end}}{{end}}{{if .VisibleFlags}}
FLAGS:{{range .Flags}}
	{{.}}{{end}}
{{end}}`

	return app
}
