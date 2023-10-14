package main

import (
	"github.com/enuesaa/walkin/internal/cli"
)

func main() {
	app := cli.CreateApp()
	app.Execute()
}
