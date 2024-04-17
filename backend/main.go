package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"go-wallet-api/configs"
	"go-wallet-api/database"
	"go-wallet-api/database/seeders"
	"go-wallet-api/internal"
	"go-wallet-api/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func main() {
	internal.InitializeDependencies()
	// migrate tables
	database.MigrateEntities()

	// seed data
	seeders.SeedEntityData()

	app := fiber.New()
	app.Use(limiter.New(configs.AddRateLimiter())) //set rate limiter
	app.Use(cors.New(configs.AddCORS()))           //set CORS
	app.Use(logger.New())

	api := app.Group("/api")

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(200).SendString("E-Wallet API")
	})

	//register routes
	routes.NewUserRoutes(api)
	routes.RegisterAuthRoutes(api)
	routes.RegisterWalletRoutes(api)

	log.Fatal(app.Listen(fmt.Sprintf(":%s", os.Getenv("PORT"))))
}
