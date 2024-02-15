package routers

import (
	"api/api/controller"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("", controller.HomePage)
}

