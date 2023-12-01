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
	authRoutes := api.Group("/auth")
	usersRoutes := api.Group("/users").Use(config.AddAuthentication())     // set authentication
	walletsRoutes := api.Group("/wallets").Use(config.AddAuthentication()) // set authentication

	models.InitValidation() //enable validation
	config.ConnectDatabase()
	config.RegisterRepositories(config.DbCtx)
	config.RegisterServices()

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(200).SendString("E-Wallet API")
	})

	authRoutes.Post("/login", middlewares.ValidateLoginRequest, handlers.AuthenticateUserHandler)

	//Users routes
	usersRoutes.Get("/", middlewares.RequiresAuthorization([]string{models.ADMIN_ROLE}), handlers.GetAllUsersHandler)
	usersRoutes.Get("/:id", middlewares.RequiresAuthorization([]string{models.ADMIN_ROLE, models.USER_ROLE}), handlers.GetUserByIdHandler)
	usersRoutes.Post("/", middlewares.RequiresAuthorization([]string{models.ADMIN_ROLE, models.USER_ROLE}), middlewares.ValidateUserRequest, handlers.CreateUserHandler)
	usersRoutes.Put("/:id", middlewares.RequiresAuthorization([]string{models.ADMIN_ROLE, models.USER_ROLE}), middlewares.ValidateUserRequest, handlers.UpdateUserHandler)
	usersRoutes.Delete("/:id", middlewares.RequiresAuthorization([]string{models.ADMIN_ROLE, models.USER_ROLE}), handlers.DeleteUserHandler)

	//Wallets routes
	walletsRoutes.Get("/", middlewares.RequiresAuthorization([]string{models.ADMIN_ROLE, models.USER_ROLE}), handlers.GetAllWalletsHandler)
	walletsRoutes.Get("/:id", middlewares.RequiresAuthorization([]string{models.ADMIN_ROLE, models.USER_ROLE}), handlers.GetWalletByIdHandler)
	walletsRoutes.Post("/", middlewares.RequiresAuthorization([]string{models.ADMIN_ROLE, models.USER_ROLE}), middlewares.ValidateWalletRequest, handlers.CreateWalletHandler)
	walletsRoutes.Put("/:id", middlewares.RequiresAuthorization([]string{models.ADMIN_ROLE, models.USER_ROLE}), middlewares.ValidateWalletRequest, handlers.UpdateWalletHandler)
	walletsRoutes.Delete("/:id", middlewares.RequiresAuthorization([]string{models.ADMIN_ROLE, models.USER_ROLE}), handlers.DeleteWalletHandler)

	log.Fatal(app.Listen(fmt.Sprintf(":%s", os.Getenv("PORT"))))
}
