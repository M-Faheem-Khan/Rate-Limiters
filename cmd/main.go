package main

import (
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/m-faheem-khan/rate-limiter/pkg/database"
	"github.com/m-faheem-khan/rate-limiter/pkg/middlewares"
)

func main() {
	app := fiber.New()
	rdb := database.RedisConnection()

	app.Use(func(c fiber.Ctx) error {
		return middlewares.TokenBucket(&middlewares.Config{
			MAX_REQUESTS: 10,
			TIME:         time.Second * 10,
			REDIS:        rdb,
		}, c)
	})

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Continue Making Requests")
	})

	app.Listen(":3000")
}

// EOF
