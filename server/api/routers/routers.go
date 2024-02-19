package routers

import (
	"api/api/controller"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/", controller.Home)
	app.Get("/api/clients", controller.Clients)
	app.Post("/api/register", controller.Register)
	app.Post("/api/login", controller.Login)
}
