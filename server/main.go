package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()

	app.Get("/user", func(c fiber.Ctx) error {
		return c.SendString("Hello world")
	})

	log.Fatal(app.Listen(":4000"))
}
