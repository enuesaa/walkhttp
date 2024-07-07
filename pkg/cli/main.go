package cli

import (
	"fmt"
	"time"

	"github.com/enuesaa/walkhttp/pkg/invoke"
	"github.com/enuesaa/walkhttp/pkg/repository"
	"github.com/pkg/browser"
	"github.com/urfave/cli/v2"
)

func New(repos repository.Repos) *cli.App {
	app := &cli.App{
		Name:    "walkhttp",
		Version: "0.0.9",
		Usage:   "A CLI tool to call http endpoint with browser or prompt.",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:     "port",
				Usage:    "Serve port",
				Value:    3000,
				Category: "SERVE",
			},
		},
		Commands: []*cli.Command{
			NewGetCommand(repos),
			NewPostCommand(repos),
			NewPutCommand(repos),
			NewDeleteCommand(repos),
			NewOptionsCommand(repos),
		},
		Args:      true,
		ArgsUsage: "commands",
		Action: func(c *cli.Context) error {
			port := c.Int("port")

			repos.Log.Printf("\n")
			repos.Log.Printf("Serving web console on localhost:%d.\n", port)
			repos.Log.Printf("\n")

			go func() {
				time.Sleep(1 * time.Second)
				url := fmt.Sprintf("http://localhost:%d", port)
				browser.OpenURL(url)
			}()

			return invoke.Serve(repos, port)

		},
		Suggest: true,
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
FLAGS:{{range .VisibleFlagCategories}}{{if len .Name}}  {{.Name}}{{end}}
	{{range .Flags}}{{.}}
	{{end}}
{{end}}{{end}}`

	return app
}
