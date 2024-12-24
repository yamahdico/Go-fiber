package main

import (
	"fmt"
	"log"
	"os"

	db "fiberProject/Fiber3/database"

	"fiberProject/Fiber3/seeders"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// Initialize Fiber
	app := fiber.New()

	// Load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Warning: .env file not found. Defaulting to environment variables.")
	}

	//DB stuff
	db.Connect()
	db.AutoMigrate()
	seeders.SeedUsers()

	// Set up a simple GET route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(&fiber.Map{
			"message": "Everything is working fine",
		})
	})

	// Start the Fiber app
	port := os.Getenv("PORT")
	if port == "" {
		port = "3001" // Default to port 3000 if PORT is not set
	}
	log.Fatal(app.Listen(fmt.Sprintf(":%s", port)))
}
