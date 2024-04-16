package configs

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func AddCORS() cors.Config {
	return cors.Config{
		Next:             nil,
		AllowOriginsFunc: nil,
		AllowOrigins:     "*",
		AllowMethods: strings.Join([]string{
			fiber.MethodGet,
			fiber.MethodPost,
			fiber.MethodHead,
			fiber.MethodPut,
			fiber.MethodDelete,
			//fiber.MethodPatch,
		}, ","),
		AllowHeaders:     "",
		AllowCredentials: false,
		ExposeHeaders:    "",
		MaxAge:           0,
	}
}
