package validations

import (
	"go-wallet-api/features/shared/di"
	shared "go-wallet-api/features/shared/models"
	"go-wallet-api/features/shared/utils"
	"go-wallet-api/features/wallet/business_logic/models"

	"github.com/gofiber/fiber/v2"
)

// Validates wallet request
func ValidateWalletRequest(c *fiber.Ctx) error {
	body := new(models.WalletRequest)
	c.BodyParser(&body)

	validator := di.WithSharedInjector.Validator
	err := validator.Struct(body)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(shared.APIResponse{
			Code:    422,
			Message: "Validation Errors",
			Errors:  utils.GetValidationErrors(err, models.WalletRequest{}),
		})
	}
	return c.Next()
}
