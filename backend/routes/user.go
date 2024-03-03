package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-wallet-api/config"
	"go-wallet-api/features/users/business_logic/handlers"
	"go-wallet-api/middlewares"
	"go-wallet-api/models"
)

func NewUserRoutes(api fiber.Router) {
	usersRoutes := api.Group("/users").Use(config.AddAuthentication()) // set authentication
	//Users routes
	usersRoutes.Get("/", middlewares.RequiresAuthorization([]string{models.ADMIN_ROLE}), handlers.GetAllUsersHandler)
	usersRoutes.Get("/:id", middlewares.RequiresAuthorization([]string{models.ADMIN_ROLE, models.USER_ROLE}), handlers.GetUserByIdHandler)
	usersRoutes.Post("/", middlewares.RequiresAuthorization([]string{models.ADMIN_ROLE, models.USER_ROLE}), middlewares.ValidateUserRequest, handlers.CreateUserHandler)
	usersRoutes.Put("/:id", middlewares.RequiresAuthorization([]string{models.ADMIN_ROLE, models.USER_ROLE}), middlewares.ValidateUserRequest, handlers.UpdateUserHandler)
	usersRoutes.Delete("/:id", middlewares.RequiresAuthorization([]string{models.ADMIN_ROLE, models.USER_ROLE}), handlers.DeleteUserHandler)
}
