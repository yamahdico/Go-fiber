package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbClient *gorm.DB

func main() {
	// Initialize Fiber
	app := fiber.New()

	// Load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Warning: .env file not found. Defaulting to environment variables.")
	}

	// Connect to Database
	dsn := os.Getenv("DB_DSN")
	//dsn := "user:password@tcp(127.0.0.1:3306)/database?charset=utf8mb4&parseTime=True&loc=Local"
	dbClient, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}
	log.Println("Database connected")

	// Get the underlying SQL database connection and defer its closure
	sqlDB, err := dbClient.DB()
	if err != nil {
		log.Fatal("Error getting database connection: ", err)
	}
	defer sqlDB.Close()

	// Set up a simple GET route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(&fiber.Map{
			"message": "Everything is working fine",
		})
	})

	// Start the Fiber app
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000" // Default to port 3000 if PORT is not set
	}
	log.Fatal(app.Listen(fmt.Sprintf(":%s", port)))
}
