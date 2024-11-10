package command

import (
	"fmt"
	"time"

	"github.com/enuesaa/walkhttp/internal/command/prompt"
	"github.com/enuesaa/walkhttp/internal/repository"
	"github.com/enuesaa/walkhttp/internal/router"
	"github.com/pkg/browser"
	"github.com/urfave/cli/v2"
)

func New(repos repository.Repos) *cli.App {
	app := &cli.App{
		Name:    "walkhttp",
		Version: "0.0.10",
		Usage:   "A CLI tool to call http endpoint with browser or prompt.",
		Commands: []*cli.Command{
			NewGetCommand(repos),
			NewPostCommand(repos),
			NewPutCommand(repos),
			NewDeleteCommand(repos),
			NewOptionsCommand(repos),
		},
		Suggest: true,
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:  "port",
				Usage: "Serve port",
				Value: 3000,
			},
		},
		Action: func(c *cli.Context) error {
			port := c.Int("port")
			prompt.PrintBanner(repos)

			go func() {
				time.Sleep(1 * time.Second)
				url := fmt.Sprintf("http://localhost:%d", port)
				browser.OpenURL(url)
			}()

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
