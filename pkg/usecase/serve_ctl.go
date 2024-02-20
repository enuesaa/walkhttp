package usecase

import (
	"github.com/enuesaa/walkin/ctlweb"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Serve() error {
	app := fiber.New()
	app.Use(logger.New())

	serveCtl := NewServeCtl()
	app.Post("/api/invoke", serveCtl.CreateInvoke)
	app.Get("/*", ctlweb.Serve)

	if err := app.Listen(":3000"); err != nil {
		return err
	}
	return nil
}