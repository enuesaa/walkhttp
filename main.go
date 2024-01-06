package main

import (
	"log"

	"github.com/enuesaa/walkin/pkg/cli"
	"github.com/enuesaa/walkin/pkg/repository"
	"github.com/spf13/cobra"
)

func init() {
	log.SetFlags(0)
}

func main() {
	app := &cobra.Command{
		Use:     "walkin",
		Short:   "A CLI Tool to debug web api",
		Version: "0.0.2",
	}

	repos := repository.NewRepos()
	app.AddCommand(cli.CreateListCmd(repos))
	app.AddCommand(cli.CreateTransformRequestCmd(repos))
	app.AddCommand(cli.CreateInvokeCmd(repos))
	app.AddCommand(cli.CreateServeCmd(repos))
	app.AddCommand(cli.CreateNewCmd(repos))

	// disable default
	app.SetHelpCommand(&cobra.Command{Hidden: true})
	app.CompletionOptions.DisableDefaultCmd = true
	app.SilenceUsage = true
	app.PersistentFlags().SortFlags = false
	app.PersistentFlags().BoolP("help", "", false, "Show help information")
	app.PersistentFlags().BoolP("version", "", false, "Show version")

	app.Execute()
}
