package main

import (
	_ "embed"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/enuesaa/walkin/controller"
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

func main() {
	workdir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("running..\n\n")
	fmt.Printf("workdir: %s\n", workdir)

	app := fiber.New(fiber.Config{
		ErrorHandler: errorHandler,
	})
	app.Use(recover.New())

	queryCtl := controller.QueryController{}
	app.Post("/query", queryCtl.Query)

	app.Listen(":3000")
}
