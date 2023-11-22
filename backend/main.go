package main

import (
	"fmt"
	"go-wallet-api/config"
	"go-wallet-api/handlers"
	"go-wallet-api/middlewares"
	"go-wallet-api/models"
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

	api := app.Group("/api")
	usersRoutes := api.Group("/users")
	walletsRoutes := api.Group("/wallets")

	models.InitValidation()
	config.ConnectDatabase()

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(200).SendString("E-Wallet API")
	})

	//Users routes
	usersRoutes.Get("/", handlers.GetAllUsersHandler)
	usersRoutes.Get("/:id", handlers.GetUserByIdHandler)
	usersRoutes.Post("/", middlewares.ValidateUserRequest, handlers.CreateUserHandler)
	usersRoutes.Put("/:id", middlewares.ValidateUserRequest, handlers.UpdateUserHandler)

	//Wallets routes
	walletsRoutes.Get("/", handlers.GetAllWalletsHandler)
	walletsRoutes.Get("/:id", handlers.GetWalletByIdHandler)
	walletsRoutes.Post("/", middlewares.ValidateWalletRequest, handlers.CreateWalletHandler)
	walletsRoutes.Put("/:id", middlewares.ValidateWalletRequest, handlers.UpdateWalletHandler)

	log.Fatal(app.Listen(fmt.Sprintf(":%s", os.Getenv("PORT"))))
}
