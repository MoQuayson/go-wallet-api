package middlewares

import (
	"go-wallet-api/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt/v5"
)

// Gets user claims/ information from jwt token
func GetUserClaims(ctx *fiber.Ctx) models.User {
	data := ctx.Locals("user").(*jwt.Token)
	claims := data.Claims.(jwt.MapClaims)

	if claims == nil {
		return models.User{}
	}

	return models.User{
		ID:       uuid.FromStringOrNil(claims["id"].(string)),
		Name:     models.GetStringFromInterface(claims["name"]),
		Email:    models.GetStringFromInterface(claims["email"]),
		PhoneNum: models.ConvertToNullString(claims["phone_num"]),
		Role:     models.GetStringFromInterface(claims["role"]),
	}
}
