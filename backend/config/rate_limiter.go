package config

import (
	"go-wallet-api/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

/*
This is where you set the rate limirer for the API
*/
func AddRateLimiter() limiter.Config {
	return limiter.Config{
		Max:        30,
		Expiration: 1 * time.Minute,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP()
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(&models.APIResponse{
				Code:    fiber.StatusTooManyRequests,
				Message: "Too Many Requests",
			})
		},
		SkipFailedRequests:     false,
		SkipSuccessfulRequests: false,
		LimiterMiddleware:      limiter.SlidingWindow{},
	}
}
