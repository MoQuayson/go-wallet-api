package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-wallet-api/config"
	"go-wallet-api/features/shared/utils/enums"
	"go-wallet-api/features/shared/utils/validations"
	"go-wallet-api/features/users/business_logic/handlers"
	"go-wallet-api/middlewares"
)

func NewUserRoutes(api fiber.Router) {
	usersRoutes := api.Group("/users").Use(config.AddAuthentication()) // set authentication
	//Users routes
	usersRoutes.Get("/", middlewares.RequiresAuthorization([]string{enums.RoleType_Admin}), handlers.GetAllUsersHandler)
	usersRoutes.Get("/:id", middlewares.RequiresAuthorization([]string{enums.RoleType_Admin, enums.RoleType_User}), handlers.GetUserByIdHandler)
	usersRoutes.Post("/", middlewares.RequiresAuthorization([]string{enums.RoleType_Admin, enums.RoleType_User}), validations.ValidateUserRequest, handlers.CreateUserHandler)
	usersRoutes.Put("/:id", middlewares.RequiresAuthorization([]string{enums.RoleType_Admin, enums.RoleType_User}), validations.ValidateUserRequest, handlers.UpdateUserHandler)
	usersRoutes.Delete("/:id", middlewares.RequiresAuthorization([]string{enums.RoleType_Admin, enums.RoleType_User}), handlers.DeleteUserHandler)
}
