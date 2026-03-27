package middelware

import (
	ratelimiter "github.com/ShubhamKharde45/rate_limiter/internal/infrastructure/rate_limiter"
	"github.com/gofiber/fiber/v3"
)

func HandleRequestRate(rateLimiter *ratelimiter.RateLimiter) fiber.Handler {

	return func(c fiber.Ctx) error {
		IP := c.IP()

		if allowed := rateLimiter.IsAllowed(IP); allowed {
			return c.Next()
		}

		return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
			"error": "Too many requests.",
		})
	}

}
