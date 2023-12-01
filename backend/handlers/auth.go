package handlers

import (
	"go-wallet-api/models"
	"go-wallet-api/requests"
	"go-wallet-api/services"

	"github.com/gofiber/fiber/v2"
)

// Login endpoint
func AuthenticateUserHandler(ctx *fiber.Ctx) error {
	payload := requests.LoginRequest{}
	//get query data
	q := ctx.Queries()
	email := q["email"]
	password := q["password"]

	if email != "" {
		payload.Email = email
		payload.Password = password
	} else {
		//get body data
		if err := ctx.BodyParser(&payload); err != nil {
			return ctx.Status(500).JSON(&models.APIResponse{
				Code:    500,
				Message: models.AUTHENTICATE_USER_ERR,
			})
		}
	}

	//get service and authenticate user
	s := services.GetAuthService()
	isSuccess, user := s.AuthenticateUser(payload)

	if !isSuccess {
		return ctx.Status(404).JSON(&models.APIResponse{
			Code:    404,
			Message: models.INVALID_CREDENTIALS,
		})
	}

	return ctx.Status(200).JSON(&models.APIResponse{
		Code:    200,
		Message: models.VALID_CREDENTIALS,
		Data:    user,
	})
}
