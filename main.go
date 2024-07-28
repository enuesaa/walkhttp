package main

import (
	"log"
	"os"

	"github.com/enuesaa/walkhttp/pkg/command"
	"github.com/enuesaa/walkhttp/pkg/repository"
	"github.com/urfave/cli/v2"
)

func main() {
	repos := repository.New()

	app := cli.App{
		Name:    "walkhttp",
		Version: "0.0.9",
		Usage:   "A CLI tool to call http endpoint with browser or prompt.",
		Commands: []*cli.Command{
			command.NewGetCommand(repos),
			command.NewPostCommand(repos),
			command.NewPutCommand(repos),
			command.NewDeleteCommand(repos),
			command.NewOptionsCommand(repos),
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
FLAGS:{{range .Flags}}
	{{.}}{{end}}
{{end}}`

	if err := app.Run(os.Args); err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
}
