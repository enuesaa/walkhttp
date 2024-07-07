package main

import (
	"log"
	"os"

	"github.com/enuesaa/walkhttp/pkg/command"
	"github.com/enuesaa/walkhttp/pkg/repository"
)

func main() {
	repos := repository.New()
	app := command.New(repos)

	if err := app.Run(os.Args); err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
}
