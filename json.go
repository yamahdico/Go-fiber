package main

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"message": "Hello",
			"status":  http.StatusOK,
		})
	})

	err := app.Listen(":2334")

	if err != nil {
		log.Fatal(err)
	}
}
