package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/ShubhamKharde45/rate_limiter/internal/domain"
	"github.com/ShubhamKharde45/rate_limiter/internal/infrastructure/cache"
	ratelimiter "github.com/ShubhamKharde45/rate_limiter/internal/infrastructure/rate_limiter"
	"github.com/ShubhamKharde45/rate_limiter/internal/transport/http/middelware"
	"github.com/gofiber/fiber/v3"
)

var mu sync.Mutex
var wg sync.WaitGroup

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	store := cache.NewRedisCache[*domain.Bucket]("Shubham@100")
	rateLimiter := ratelimiter.NewRateLimiter(store, &mu)
	fmt.Println("Rate limiter started...")

	app := fiber.New()

	app.Use(middelware.HandleRequestRate(rateLimiter))

	app.Get("/", func(c fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"message": "Hello from Server : " + port,
		})
	})

	fmt.Printf("Server started at %s", port)
	err := app.Listen(":" + port)

	if err != nil {
		panic(err)
	}

}
