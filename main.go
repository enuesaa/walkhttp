package main

import (
	"fmt"
	"log"
	"os"
	_ "embed"

	"github.com/enuesaa/walkin/controller"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

//go:embed tmp/dist/index.html
var indexHtmlBytes []byte
//go:embed tmp/dist/index.js
var indexJsBytes []byte

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
	app.Get("/", func(c *fiber.Ctx) error {
		c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
		return c.SendString(string(indexHtmlBytes))
	})
	app.Get("/index.js", func(c *fiber.Ctx) error {
		c.Set(fiber.HeaderContentType, fiber.MIMETextJavaScript)
		return c.SendString(string(indexJsBytes))
	})

	app.Listen(":3000")
}
