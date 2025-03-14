package main

import (
	"log"
	"os"

	"github.com/enuesaa/walkhttp/internal/command"
	"github.com/enuesaa/walkhttp/internal/repository"
)

func main() {
	repos, err := repository.New()
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
	app := command.New(repos)

	if err := app.Run(os.Args); err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
}
