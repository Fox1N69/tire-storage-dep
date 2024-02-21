package main

import (
	"api/api/routers"
	"log"

	"api/api/database"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

func main() {
	database.Connection()
	app := fiber.New()

	routers.Setup(app)

	log.Fatal(app.Listen(":4000"))
}
