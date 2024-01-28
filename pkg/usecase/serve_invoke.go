package usecase

import (
	"github.com/gofiber/fiber/v2"
)

type CreateInvokeResponseData struct {}
func (ctl *ServeCtl) CreateInvoke(c *fiber.Ctx) error {
	res := CreateInvokeResponseData{}

	return c.JSON(res)
}
