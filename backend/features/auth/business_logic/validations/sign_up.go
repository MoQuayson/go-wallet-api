package validations

import (
	"github.com/gofiber/fiber/v2"
	"go-wallet-api/features/auth/business_logic/requests"
	"go-wallet-api/features/shared/di"
	shared "go-wallet-api/features/shared/models"
	"go-wallet-api/features/shared/utils"
)

func ValidateSignUpRequest(c *fiber.Ctx) error {
	req := new(requests.SignUpRequest)
	_ = c.BodyParser(&req)

	validator := di.WithSharedInjector.Validator
	err := validator.Struct(req)

	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(shared.APIResponse{
			Code:    422,
			Message: "Validation Errors",
			Errors:  utils.GetValidationErrors(err, requests.SignUpRequest{}),
		})
	}
	return c.Next()
}
