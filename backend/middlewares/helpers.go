package middlewares

import (
	userModel "go-wallet-api/features/users/business_logic/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt/v5"
)

// GetUserClaims Gets user claims information from jwt token
func GetUserClaims(ctx *fiber.Ctx) *userModel.User {
	data := ctx.Locals("user").(*jwt.Token)
	claims := data.Claims.(jwt.MapClaims)

	if claims == nil {
		return nil
	}

	phoneNumClaim := getStringFromInterface(claims["phone_num"])

	return &userModel.User{
		ID:       uuid.FromStringOrNil(claims["id"].(string)),
		Name:     getStringFromInterface(claims["name"]),
		Email:    getStringFromInterface(claims["email"]),
		PhoneNum: &phoneNumClaim,
		Role:     getStringFromInterface(claims["role"]),
	}
}

func getStringFromInterface(v interface{}) string {
	if v == nil {
		return ""
	}

	return v.(string)
}
