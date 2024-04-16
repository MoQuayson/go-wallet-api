package utils

import (
	"github.com/gofiber/fiber/v2"
	shared "go-wallet-api/features/shared/models"
	"go-wallet-api/features/shared/utils/enums"
)

func SetErrorResponseMessage(ctx *fiber.Ctx, msg enums.ResponseMsg) error {
	return ctx.Status(500).JSON(&shared.APIResponse{
		Code:    500,
		Message: msg,
	})
}
