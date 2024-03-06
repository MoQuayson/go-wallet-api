package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"go-wallet-api/config"
	"go-wallet-api/features/shared/utils/validations"
	"go-wallet-api/internal"
	"go-wallet-api/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func main() {
	app := fiber.New()
	app.Use(limiter.New(config.AddRateLimiter())) //set rate limiter
	app.Use(cors.New(config.AddCORS()))           //set CORS
	app.Use(logger.New())

	api := app.Group("/api")

	validations.InitValidation() //enable validation
	config.ConnectDatabase()
	internal.InitializeDependencies()
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(200).SendString("E-Wallet API")
	})

	//register routes
	routes.NewUserRoutes(api)
	routes.RegisterAuthRoutes(api)
	routes.RegisterWalletRoutes(api)

	log.Fatal(app.Listen(fmt.Sprintf(":%s", os.Getenv("PORT"))))
}
