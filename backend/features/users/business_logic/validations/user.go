package validations

import (
	"github.com/gofiber/fiber/v2"
	"go-wallet-api/features/shared/di"
	shared "go-wallet-api/features/shared/models"
	"go-wallet-api/features/shared/utils"
	"go-wallet-api/features/users/business_logic/requests"
)

func ValidateUserRequest(c *fiber.Ctx) error {
	body := new(requests.UserRequest)
	c.BodyParser(&body)

	validator := di.WithSharedInjector.Validator
	err := validator.Struct(body)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(shared.APIResponse{
			Code:    422,
			Message: "Validation Errors",
			Errors:  utils.GetValidationErrors(err, requests.UserRequest{}),
		})
	}
	return c.Next()
}
