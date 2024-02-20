package main

import (
	"github.com/enuesaa/walkin/pkg/cli"
	"github.com/enuesaa/walkin/pkg/repository"
	"github.com/spf13/cobra"
)

func main() {
	app := &cobra.Command{
		Use:     "walkin",
		Short:   "http utility tool",
		Version: "0.0.2",
	}

	repos := repository.NewRepos()
	// app.AddCommand(cli.CreateServeCmd(repos))
	// app.AddCommand(cli.CreateCtlCmd(repos))
	// app.AddCommand(cli.CreateHistoryCmd(repos))
	app.AddCommand(cli.CreateGetCmd(repos))
	app.AddCommand(cli.CreatePostCmd(repos))
	app.AddCommand(cli.CreatePutCmd(repos))
	app.AddCommand(cli.CreateDeleteCmd(repos))
	app.AddCommand(cli.CreateOptionsCmd(repos))
	app.AddCommand(cli.CreateReuseCmd(repos))

	// disable default
	app.SetHelpCommand(&cobra.Command{Hidden: true})
	app.CompletionOptions.DisableDefaultCmd = true
	app.SilenceUsage = true
	app.PersistentFlags().SortFlags = false
	app.PersistentFlags().BoolP("help", "", false, "Show help information")
	app.PersistentFlags().BoolP("version", "", false, "Show version")

	app.Execute()
}
