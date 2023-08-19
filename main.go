package main

import (
	"flag"
    "fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
    flag.Parse()
    args := flag.Args()
	if len(args) == 0 {
		fmt.Println("Required argument missing: WORKDIR")
		return
	}

	workdir := args[0]
	fmt.Println(workdir)

    app := fiber.New()
	
	api := app.Group("/api")
	// create api routes here

    app.Get("/", func(c *fiber.Ctx) error {
		// should return html or assets
        return c.SendString("")
    })
    app.Listen(":3000")
}
