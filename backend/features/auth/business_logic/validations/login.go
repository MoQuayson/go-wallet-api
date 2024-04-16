package validations

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"go-wallet-api/features/auth/business_logic/requests"
	"go-wallet-api/features/shared/di"
	shared "go-wallet-api/features/shared/models"
	"go-wallet-api/features/shared/utils"
)

// ValidateLoginRequest Validates user requests
func ValidateLoginRequest(c *fiber.Ctx) error {
	var payload requests.LoginRequest

	if c.Query("email") != "" {
		log.Info("Endpoint has queries")
		payload.Email = c.Query("email")
		payload.Password = c.Query("password")
	} else {
		log.Info("Endpoint has Body")
		_ = c.BodyParser(&payload)
	}

	validator := di.WithSharedInjector.Validator

	err := validator.Struct(payload)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(shared.APIResponse{
			Code:    422,
			Message: "Validation Errors",
			Errors:  utils.GetValidationErrors(err, requests.LoginRequest{}),
		})
	}

	return c.Next()
}
