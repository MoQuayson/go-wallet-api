package validations

import (
	"github.com/gofiber/fiber/v2"
	shared "go-wallet-api/features/shared/models"
	userModel "go-wallet-api/features/users/business_logic/app/models"
)

func ValidateUserRequest(c *fiber.Ctx) error {
	body := new(userModel.UserRequest)
	c.BodyParser(&body)

	err := Validator.Struct(body)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(shared.APIResponse{
			Code:    422,
			Message: "Validation Errors",
			Errors:  GetValidationErrors(err, userModel.UserRequest{}),
		})
	}
	return c.Next()
}
