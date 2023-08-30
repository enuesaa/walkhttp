package main

import (
	"fmt"
	"log"
	"os"

	"github.com/enuesaa/walkin/controller"
	"github.com/gofiber/fiber/v2"
)

func main() {
	workdir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("running..\n\n")
	fmt.Printf("workdir: %s\n", workdir)

	app := fiber.New()

	queryCtl := controller.QueryController{}
	app.Post("/query", queryCtl.Query)

	app.Listen(":3000")
}
