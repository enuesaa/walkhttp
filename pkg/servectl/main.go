package servectl

import (
	"github.com/enuesaa/walkin/ctlweb"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func New() ServeCtl {
	return ServeCtl{}
}
type ServeCtl struct {}

func (srv *ServeCtl) Run() error {
	app := fiber.New()
	app.Use(logger.New())

	app.Post("/api/invoke", srv.CreateInvoke)
	app.Get("/*", ctlweb.Serve)

	return app.Listen(":3000")
}