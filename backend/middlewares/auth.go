package middlewares

import (
	"go-wallet-api/models"
	"go-wallet-api/requests"
	"go-wallet-api/services"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

// Validates user request
func ValidateLoginRequest(c *fiber.Ctx) error {
	var payload requests.LoginRequest

	if c.Query("email") != "" {
		log.Info("Endpoint has queries")
		payload.Email = c.Query("email")
		payload.Password = c.Query("password")
	} else {
		log.Info("Endpoint has Body")
		c.BodyParser(&payload)
	}

	err := models.Validator.Struct(payload)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(models.APIResponse{
			Code:    422,
			Message: "Validation Errors",
			Errors:  models.GetValidationErrors(err, requests.LoginRequest{}),
		})
	}

	return c.Next()
}

// Middleware to check if user is authenticated
func RequiresAuthentication() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   jwtware.SigningKey{Key: []byte(services.GetJWTSecret())},
		ErrorHandler: JwtError,
	})
}

func JwtError(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusUnauthorized).JSON(models.APIResponse{
		Code:    fiber.StatusUnauthorized,
		Message: models.UNAUTHENTICATED_USER,
	})
}

// Checks if user has the correct authorization
func RequiresAuthorization(roles []string) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		user := GetUserClaims(ctx)

		for _, r := range roles {
			if r == user.Role {
				return ctx.Next()
			}
		}

		return ctx.Status(fiber.StatusUnauthorized).JSON(models.APIResponse{
			Code:    fiber.StatusUnauthorized,
			Message: models.UNAUTHORIZED_USER,
		})
	}

}
