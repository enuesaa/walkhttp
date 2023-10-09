package handler

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type ErrorResponse struct {
	Error string `json:"error"`
}
func errorHandler(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}
	return ctx.Status(code).JSON(ErrorResponse { Error: "Internal Sever Error" })
}

func Handle(port string) {
	workdir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("running on %s\n", workdir)
		
	app := fiber.New(fiber.Config{
		ErrorHandler: errorHandler,
	})
	app.Use(recover.New())

	queryCtl := QueryController{}
	app.Post("/query", queryCtl.Query)

	app.Get("/", func(c *fiber.Ctx) error {
		f, err := os.Open("./testdata/index.html")
		if err != nil {
			return err
		}
		defer f.Close()

		content, err := io.ReadAll(f)
		if err != nil {
			return err
		}
		c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)

		return c.SendString(string(content))
	})

	addr := fmt.Sprintf(":%s", port)
	app.Listen(addr)
}