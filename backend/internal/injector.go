package internal

import (
	"go-wallet-api/config"
	userDI "go-wallet-api/features/users/di"
	"log"
)

func InitializeDependencies() {

	if userDI.WithUserInjector = userDI.NewUserInjector(config.DbCtx); userDI.WithUserInjector == nil {
		log.Println("could not initialize user injector")
	}
}
