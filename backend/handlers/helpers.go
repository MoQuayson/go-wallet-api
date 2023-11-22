package handlers

import (
	"go-wallet-api/config"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetDbConnection(c *fiber.Ctx) *gorm.DB {
	return config.DbCtx
}
