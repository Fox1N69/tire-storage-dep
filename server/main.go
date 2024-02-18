package main

import (
	"api/api/database"
	"api/api/routers"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	routers.Setup(app)

	log.Fatal(app.Listen(":4000"))

	database.Connection() 
}
