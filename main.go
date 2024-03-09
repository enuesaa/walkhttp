package main

import (
	"github.com/enuesaa/walkin/pkg/cli"
	"github.com/enuesaa/walkin/pkg/repository"
	"github.com/spf13/cobra"
)

func main() {
	app := &cobra.Command{
		Use:     "walkin",
		Short:   "A CLI tool to serve local api gateway",
		Version: "0.0.2",
	}

	repos := repository.NewRepos()
	app.AddCommand(cli.CreateInitCmd(repos))

	// http commands
	// read port from config file
	// call http api from browser or server (golang)
	app.AddCommand(cli.CreateGetCmd(repos))
	app.AddCommand(cli.CreatePostCmd(repos))
	app.AddCommand(cli.CreatePutCmd(repos))
	app.AddCommand(cli.CreateDeleteCmd(repos))
	app.AddCommand(cli.CreateOptionsCmd(repos))
	app.AddCommand(cli.CreateUpCmd(repos))

	// disable default
	app.SetHelpCommand(&cobra.Command{Hidden: true})
	app.CompletionOptions.DisableDefaultCmd = true
	app.SilenceUsage = true
	app.PersistentFlags().SortFlags = false
	app.PersistentFlags().BoolP("help", "", false, "Show help information")
	app.PersistentFlags().BoolP("version", "", false, "Show version")

	app.Execute()
}
