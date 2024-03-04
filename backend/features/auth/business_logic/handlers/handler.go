package handlers

import (
	"go-wallet-api/features/auth/business_logic/app/models"
	authDI "go-wallet-api/features/auth/di"
	shared "go-wallet-api/features/shared/models"
	"go-wallet-api/features/shared/utils/enums"
	"log"

	"github.com/gofiber/fiber/v2"
)

func AuthenticateUserHandler(ctx *fiber.Ctx) error {
	req := &models.LoginRequest{}
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
			return returnErrorResponse(ctx, enums.ResponseMsg_AuthenticateUserErr)
		}
	}

	srv := authDI.WithAuthInjector.Service
	user, err := srv.AuthenticateUser(req)

	if err != nil {
		log.Println(err)
		return returnErrorResponse(ctx, enums.ResponseMsg_AuthenticateUserErr)
	}

	if err != nil && user == nil {
		log.Println(err)
		return returnErrorResponse(ctx, enums.ResponseMsg_InvalidCredentials)
	}

	return ctx.Status(200).JSON(&shared.APIResponse{
		Code:    200,
		Message: enums.ResponseMsg_ValidCredentials,
		Data:    user,
	})
}
