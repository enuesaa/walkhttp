package command

import (
	"fmt"

	"github.com/enuesaa/walkhttp/internal/command/prompt"
	"github.com/enuesaa/walkhttp/internal/repository"
	"github.com/enuesaa/walkhttp/internal/router"
	"github.com/urfave/cli/v2"
)

func New(repos repository.Repos) *cli.App {
	app := &cli.App{
		Name:    "walkhttp",
		Version: "0.0.10",
		Usage:   "A CLI tool to proxy HTTP requests for debugging.",
		Commands: []*cli.Command{
			NewGetCommand(repos),
			NewPostCommand(repos),
			NewPutCommand(repos),
			NewDeleteCommand(repos),
			NewOptionsCommand(repos),
		},
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:  "port",
				Usage: "Serve port",
				Value: 3000,
			},
			&cli.StringFlag{
				Name:  "origin",
				Usage: "Origin URL. Example: https://example.com",
			},
			&cli.IntFlag{
				Name:  "persist",
				Usage: "Persist 100 records by default",
			},
		},
		Action: func(c *cli.Context) error {
			port := c.Int("port")
			origin := c.String("origin")
			persist := c.Int("persist")
			if origin != "" {
				repos.Env.WALKHTTP_URL_BASE = origin
			}
			if persist != 0 {
				repos.Env.WALKHTTP_PERSIST_COUNT = persist
			}
			prompt.PrintBanner(repos)

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
