package di

import (
	"go-wallet-api/features/users/business_logic/app/repository"
	"go-wallet-api/features/users/business_logic/services"
	"go-wallet-api/features/users/pkg"
	"gorm.io/gorm"
)

var WithUserInjector *UserInjector

type UserInjector struct {
	DB      *gorm.DB
	Service pkg.IUserService
}

func NewUserInjector(db *gorm.DB) *UserInjector {
	injector := &UserInjector{}
	injector.Service = services.NewUserService(repository.NewUserRepository(db))
	return injector
}
