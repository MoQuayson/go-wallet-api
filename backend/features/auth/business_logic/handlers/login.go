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

func LoginUserHandler(ctx *fiber.Ctx) error {
	req := &requests.LoginRequest{}
	//get query data
	q := ctx.Queries()
	email := q["email"]
	password := q["password"]

	if email != "" {
		req.Email = email
		req.Password = password
	} else {
		//get body data
		if err := ctx.BodyParser(req); err != nil {
			log.Println(err)
			return utils.SetErrorResponseMessage(ctx, enums.ResponseMsg_AuthenticateUserErr)
		}
	}

	srv := di.WithAuthInjector.AuthService
	user, err := srv.AuthenticateUser(req)

	if err != nil {
		log.Println(err)
		return utils.SetErrorResponseMessage(ctx, enums.ResponseMsg_AuthenticateUserErr)
	}

	if err == nil && user == nil {
		log.Println(err)
		return utils.SetErrorResponseMessage(ctx, enums.ResponseMsg_InvalidCredentials)
	}

	return ctx.Status(200).JSON(&shared.APIResponse{
		Code:    200,
		Message: enums.ResponseMsg_ValidCredentials,
		Data:    user,
	})
}
