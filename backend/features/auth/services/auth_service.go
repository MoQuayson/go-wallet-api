package services

import "go-wallet-api/features/auth/pkg"

type AuthService struct {
	repo pkg.IAuthRepository
}
