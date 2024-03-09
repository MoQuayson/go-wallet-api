package pkg

import (
	"go-wallet-api/features/auth/business_logic/app/models"
	user "go-wallet-api/features/users/business_logic/app/models"
)

type IAuthService interface {
	AuthenticateUser(*models.LoginRequest) (*user.User, error)
}
