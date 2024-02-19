package controller

import (
	"github.com/gofiber/fiber/v2"
)

func Clients(c *fiber.Ctx) error {
	return c.SendString("Hello Clients page")
}
