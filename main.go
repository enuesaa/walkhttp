package main

import (
	"github.com/enuesaa/walkin/internal/cli"
)

func main() {
	// repos

	app := cli.CreateApp()
	app.Execute()
}
