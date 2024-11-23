package middleware

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

func GlobalLogger(ctx *fiber.Ctx) error {
	log.Println("This is a global logger which can be used with app.Use()")
	return ctx.Next()
}

func RequestLogger(ctx *fiber.Ctx) error {
	log.Println("The User has made a request...")
	return ctx.Next()
}

func TimeLogger(ctx *fiber.Ctx) error {
	hour, minute, second := time.Now().Clock()

	log.Printf("The time is %d:%d:%d", hour, minute, second)

	return ctx.Next()
}
