package validations

import (
	shared "go-wallet-api/features/shared/models"
	"go-wallet-api/features/wallet/business_logic/app/models"

	"github.com/gofiber/fiber/v2"
)

// Validates wallet request
func ValidateWalletRequest(c *fiber.Ctx) error {
	body := new(models.WalletRequest)
	c.BodyParser(&body)

	err := Validator.Struct(body)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(shared.APIResponse{
			Code:    422,
			Message: "Validation Errors",
			Errors:  GetValidationErrors(err, models.WalletRequest{}),
		})
	}
	return c.Next()
}
