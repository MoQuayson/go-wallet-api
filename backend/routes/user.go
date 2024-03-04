package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-wallet-api/config"
	"go-wallet-api/features/shared/utils/validations"
	"go-wallet-api/features/users/business_logic/handlers"
	"go-wallet-api/models"
)

func NewUserRoutes(api fiber.Router) {
	usersRoutes := api.Group("/users").Use(config.AddAuthentication()) // set authentication
	//Users routes
	usersRoutes.Get("/", validations.RequiresAuthorization([]string{models.ADMIN_ROLE}), handlers.GetAllUsersHandler)
	usersRoutes.Get("/:id", validations.RequiresAuthorization([]string{models.ADMIN_ROLE, models.USER_ROLE}), handlers.GetUserByIdHandler)
	usersRoutes.Post("/", validations.RequiresAuthorization([]string{models.ADMIN_ROLE, models.USER_ROLE}), validations.ValidateUserRequest, handlers.CreateUserHandler)
	usersRoutes.Put("/:id", validations.RequiresAuthorization([]string{models.ADMIN_ROLE, models.USER_ROLE}), validations.ValidateUserRequest, handlers.UpdateUserHandler)
	usersRoutes.Delete("/:id", validations.RequiresAuthorization([]string{models.ADMIN_ROLE, models.USER_ROLE}), handlers.DeleteUserHandler)
}
