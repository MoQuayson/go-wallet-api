package handlers

import (
	"github.com/gofiber/fiber/v2"
	"go-wallet-api/features/auth/business_logic/requests"
	"go-wallet-api/features/auth/di"
	shared "go-wallet-api/features/shared/models"
	"go-wallet-api/features/shared/utils"
	"go-wallet-api/features/shared/utils/enums"
	"log"
)

// RegisterUserHandler register user account
func RegisterUserHandler(ctx *fiber.Ctx) error {
	req := &requests.SignUpRequest{}
	//get body data
	if err := ctx.BodyParser(req); err != nil {
		log.Println(err)
		return utils.SetErrorResponseMessage(ctx, enums.ResponseMsg_SignUpUserErr)
	}

	srv := di.WithAuthInjector.AuthService

	// if error
	if err := srv.SignUpUser(req); err != nil {
		log.Println(err)
		return utils.SetErrorResponseMessage(ctx, enums.ResponseMsg_SignUpUserErr)
	}

	return ctx.Status(200).JSON(&shared.APIResponse{
		Code:    200,
		Message: enums.ResponseMsg_SignUpUserSucess,
	})
}
