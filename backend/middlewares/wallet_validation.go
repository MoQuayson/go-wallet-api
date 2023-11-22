package middlewares

import (
	"go-wallet-api/models"
	"go-wallet-api/requests"

	"github.com/gofiber/fiber/v2"
)

// Validates wallet request
func ValidateWalletRequest(c *fiber.Ctx) error {
	body := new(requests.WalletRequest)
	c.BodyParser(&body)

	err := models.Validator.Struct(body)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(models.APIResponse{
			Code:    422,
			Message: "Validation Errors",
			Errors:  models.GetValidationErrors(err, requests.WalletRequest{}),
		})
	}
	return c.Next()
}
