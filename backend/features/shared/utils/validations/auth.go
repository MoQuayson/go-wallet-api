package validations

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	auth "go-wallet-api/features/auth/business_logic/app/models"
	shared "go-wallet-api/features/shared/models"
)

// Validates user request
func ValidateLoginRequest(c *fiber.Ctx) error {
	var payload auth.LoginRequest

	if c.Query("email") != "" {
		log.Info("Endpoint has queries")
		payload.Email = c.Query("email")
		payload.Password = c.Query("password")
	} else {
		log.Info("Endpoint has Body")
		c.BodyParser(&payload)
	}

	err := Validator.Struct(payload)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(shared.APIResponse{
			Code:    422,
			Message: "Validation Errors",
			Errors:  GetValidationErrors(err, auth.LoginRequest{}),
		})
	}

	return c.Next()
}
