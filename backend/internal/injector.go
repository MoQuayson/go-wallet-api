package internal

import (
	"go-wallet-api/config"
	authDI "go-wallet-api/features/auth/di"
	userDI "go-wallet-api/features/users/di"
	"log"
)

func InitializeDependencies() {

	if userDI.WithUserInjector = userDI.NewUserInjector(config.DbCtx); userDI.WithUserInjector == nil {
		log.Println("could not initialize user dependencies")
	}

	if authDI.WithAuthInjector = authDI.NewAuthInjector(config.DbCtx); authDI.WithAuthInjector == nil {
		log.Println("could not initialize auth dependencies")
	}
}
