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
    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World 👋!")
    })
    app.Listen(":3000")
}
