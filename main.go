package main

import (
	"fmt"
	"log"
	"os"

	"github.com/enuesaa/walkin/controller"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	workdir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("running..\n\n")
	fmt.Printf("workdir: %s\n", workdir)

	app := fiber.New()
	// this is for develpoment
	app.Use(cors.New())

	queryCtl := controller.QueryController{}
	app.Post("/query", queryCtl.Query)

	app.Listen(":3000")
}
