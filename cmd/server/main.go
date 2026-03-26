package main

import (
	"sync"

	"github.com/ShubhamKharde45/rate_limiter/internal/delivery/middelware"
	"github.com/ShubhamKharde45/rate_limiter/internal/domain"
	cache "github.com/ShubhamKharde45/rate_limiter/internal/infrastructure/cache"
	ratelimiter "github.com/ShubhamKharde45/rate_limiter/internal/infrastructure/rate_limiter"
	"github.com/gofiber/fiber/v3"
)

var mu sync.Mutex

func main() {

	store := cache.NewMemoryStore[string, *domain.Bucket]()
	rateLimiter := ratelimiter.NewRateLimiter(store, &mu)

	app := fiber.New()

	app.Use(middelware.HandleRequestRate(rateLimiter))

	app.Get("/", func(c fiber.Ctx) error {

		dt, _ := store.Get(c.IP())
		return c.Status(200).JSON(fiber.Map{
			"message": "Success",
			"data":    &dt,
		})
	})

	app.Listen(":8080")

}
