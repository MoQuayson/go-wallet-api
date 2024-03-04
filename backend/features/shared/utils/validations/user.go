package validations

import (
	"github.com/gofiber/fiber/v2"
	userModel "go-wallet-api/features/users/business_logic/app/models"
	"go-wallet-api/models"
)

func ValidateUserRequest(c *fiber.Ctx) error {
	body := new(userModel.UserRequest)
	c.BodyParser(&body)

	err := models.Validator.Struct(body)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(models.APIResponse{
			Code:    422,
			Message: "Validation Errors",
			Errors:  models.GetValidationErrors(err, userModel.UserRequest{}),
		})
	}
	return c.Next()
}
