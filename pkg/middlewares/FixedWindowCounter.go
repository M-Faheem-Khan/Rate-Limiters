package middlewares

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v3"
)

type FixedWindowCounterConfig struct {
	REDIS        *redis.Client
	MAX_REQUESTS int
	TIME         time.Duration
}

func FixedWindowCounter(config *FixedWindowCounterConfig, c fiber.Ctx) error {
	// Checking if the IP already exists in DB
	val := config.REDIS.Get(context.Background(), c.IP())

	sVal, err := strconv.Atoi(val.Val())
	if err != nil {
		c.Status(500).SendString("Internal Error")
	}

	// First Request
	if val.Val() == "" {
		err := config.REDIS.Set(context.Background(), c.IP(), 1, config.TIME)
		if err.Val() != "OK" {
			fmt.Println("Error adding "+c.IP()+" to redis: ", err)
		}
	}

	// Too many
	if sVal >= config.MAX_REQUESTS {
		return c.Status(429).SendString("Too many requests")
	}

	// Continue making requests
	if sVal < config.MAX_REQUESTS {
		config.REDIS.Incr(context.Background(), c.IP())
	}

	return c.Next()
}

// EOF
