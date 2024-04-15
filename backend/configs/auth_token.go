package configs

import (
	"go-wallet-api/features/shared/utils"
	"go-wallet-api/middlewares"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

// Adds authentication feature using JWT Token
func AddAuthentication() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   jwtware.SigningKey{Key: []byte(utils.GetJWTSecret())},
		ErrorHandler: middlewares.JwtError,
	})
}
