package controller

import (
	"github.com/gofiber/fiber/v2"
)

type ClientS struct {
	ID          int    `json: "id"`
	FIO         string `json: "fio"`
	NumberPhone int    `json: "number_phone"`
	LastSeen    string `json: "last_seen"`
}

func Clients(c *fiber.Ctx) error {
	cln := struct {
		id   uint   `json: "id"`
		name string `json: "name"`
	}{}
	if err := c.BodyParser(&cln); err != nil {
		panic(err)
	}

	return c.JSON(cln)
}
