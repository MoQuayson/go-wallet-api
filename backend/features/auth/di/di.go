package di

import (
	"go-wallet-api/features/auth/business_logic/services"
	"go-wallet-api/features/auth/pkg"
	"go-wallet-api/features/users/business_logic/app/repository"
	"gorm.io/gorm"
)

var WithAuthInjector *AuthInjector

type AuthInjector struct {
	DB      *gorm.DB
	Service pkg.IAuthService
}

func NewAuthInjector(db *gorm.DB) *AuthInjector {
	injector := &AuthInjector{}
	injector.Service = services.NewAuthService(repository.NewUserRepository(db))
	return injector
}
