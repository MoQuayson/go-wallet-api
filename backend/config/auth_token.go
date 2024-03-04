package config

import (
	"go-wallet-api/features/shared/utils/validations"
	"go-wallet-api/services"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

// Adds authentication feature using JWT Token
func AddAuthentication() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   jwtware.SigningKey{Key: []byte(services.GetJWTSecret())},
		ErrorHandler: validations.JwtError,
	})
}
