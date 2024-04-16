package pkg

import (
	"go-wallet-api/features/auth/business_logic/requests"
	user "go-wallet-api/features/users/business_logic/models"
)

type IAuthService interface {
	AuthenticateUser(*requests.LoginRequest) (*user.User, error)
	SignUpUser(*requests.SignUpRequest) error
}
