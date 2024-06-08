package main

import (
	"github.com/enuesaa/walkin/pkg/cli"
	"github.com/enuesaa/walkin/pkg/invoke"
	"github.com/enuesaa/walkin/pkg/repository"
	"github.com/spf13/cobra"
)

func main() {
	repos := repository.NewRepos()
	app := &cobra.Command{
		Use:     "walkin",
		Short:   "A CLI tool to serve local api gateway",
		Version: "0.0.6",
		RunE: func(cmd *cobra.Command, args []string) error {
			port, _ := cmd.Flags().GetInt("port")

			return invoke.Serve(repos, port)
		},
	}
	app.Flags().Int("port", 3000, "port")

	app.AddCommand(cli.CreateGetCmd(repos))
	app.AddCommand(cli.CreatePostCmd(repos))
	app.AddCommand(cli.CreatePutCmd(repos))
	app.AddCommand(cli.CreateDeleteCmd(repos))

	// disable default
	app.SetHelpCommand(&cobra.Command{Hidden: true})
	app.CompletionOptions.DisableDefaultCmd = true
	app.SilenceUsage = true
	app.PersistentFlags().SortFlags = false
	app.PersistentFlags().BoolP("help", "", false, "Show help information")
	app.PersistentFlags().BoolP("version", "", false, "Show version")

	app.Execute()
}
