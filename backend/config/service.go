package config

import (
	"go-wallet-api/repositories"
	"go-wallet-api/services"
)

/*
This is where you register your services
This is more like a DI style
*/
func RegisterServices() {
	services.NewUserService(repositories.GetUserRepository())
	services.NewWalletService(repositories.GetWalletRepository())
	services.NewAuthService(repositories.GetUserRepository())
}
