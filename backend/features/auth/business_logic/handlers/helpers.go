package handlers

import (
	"github.com/gofiber/fiber/v2"
	shared "go-wallet-api/features/shared/models"
)

//func getUserService() *services.UserService {
//	return userDI.WithUserInjector.Service
//}

func returnErrorResponse(ctx *fiber.Ctx, msg string) error {
	return ctx.Status(500).JSON(&shared.APIResponse{
		Code:    500,
		Message: msg,
	})
}
