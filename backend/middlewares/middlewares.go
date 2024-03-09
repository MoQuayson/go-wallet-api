package middlewares

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	shared "go-wallet-api/features/shared/models"
	"go-wallet-api/features/shared/utils"
	"go-wallet-api/features/shared/utils/enums"
)

// Middleware to check if user is authenticated
func RequiresAuthentication() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   jwtware.SigningKey{Key: []byte(utils.GetJWTSecret())},
		ErrorHandler: JwtError,
	})
}

func JwtError(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusUnauthorized).JSON(shared.APIResponse{
		Code:    fiber.StatusUnauthorized,
		Message: enums.ResponseMsg_UnAuthorizedUser,
	})
}

// RequiresAuthorization Checks if user has the correct authorization
func RequiresAuthorization(roles []string) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		user := GetUserClaims(ctx)

		for _, r := range roles {
			if r == user.Role {
				return ctx.Next()
			}
		}

		return ctx.Status(fiber.StatusUnauthorized).JSON(shared.APIResponse{
			Code:    fiber.StatusUnauthorized,
			Message: enums.ResponseMsg_UnAuthorizedUser,
		})
	}

}
