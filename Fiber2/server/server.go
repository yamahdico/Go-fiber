package server

import (
	"net/http"

	"fiberProject/Fiber2/middleware"

	"github.com/gofiber/fiber/v2"
)

func Run() {
	// Initialize Fiber
	app := fiber.New()

	// Global Middlewares
	app.Use(middleware.GlobalLogger)

	// Define Routes
	app.Get("/request-logger", middleware.RequestLogger, func(ctx *fiber.Ctx) error {
		return ctx.Status(http.StatusOK).JSON(fiber.Map{
			"message": "Request log is printed out",
		})
	})
	app.Get("/time-logger", middleware.TimeLogger, func(ctx *fiber.Ctx) error {
		return ctx.Status(http.StatusOK).JSON(fiber.Map{
			"message": "Time log is printed out",
		})
	})

	app.Listen(":3003")
}
