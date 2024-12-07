package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Define a POST route
	app.Post("/api/data", func(c *fiber.Ctx) error {
		// Parse JSON body
		type RequestBody struct {
			Name  string `json:"name"`
			Email string `json:"email"`
		}

		var body RequestBody
		if err := c.BodyParser(&body); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid input",
			})
		}

		// Respond with the received data
		return c.JSON(fiber.Map{
			"message": "Data received successfully",
			"data":    body,
		})
	})

	// Start the server on port 3001
	port := "3001"
	log.Fatal(app.Listen(fmt.Sprintf(":%s", port)))

}
