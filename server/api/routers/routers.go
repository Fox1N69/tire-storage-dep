package routers

import (
	"api/api/controller"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/api/register", controller.Register)
	app.Post("/api/clients", controller.Clients)
}
