package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-wallet-api/features/auth/business_logic/handlers"
	"go-wallet-api/features/shared/utils/validations"
)

func RegisterAuthRoutes(api fiber.Router) {
	authRoutes := api.Group("/auth")
	authRoutes.Post("/login", validations.ValidateLoginRequest, handlers.AuthenticateUserHandler)
}
