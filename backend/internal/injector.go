package internal

import (
	authDI "go-wallet-api/features/auth/di"
	sharedDI "go-wallet-api/features/shared/di"
	userDI "go-wallet-api/features/users/di"
	walletDI "go-wallet-api/features/wallet/di"
	"log"
)

func InitializeDependencies() {

	if sharedDI.WithConfigsInjector = sharedDI.NewConfigsInjector(); sharedDI.WithConfigsInjector == nil {
		log.Fatalln("could not initialize configs dependencies")
	}

	if sharedDI.WithSharedInjector = sharedDI.NewSharedInjector(sharedDI.WithConfigsInjector); sharedDI.WithSharedInjector == nil {
		log.Fatalln("could not initialize shared dependencies")
	}

	if userDI.WithUserInjector = userDI.NewUserInjector(sharedDI.WithSharedInjector.DB); userDI.WithUserInjector == nil {
		log.Println("could not initialize user dependencies")
	}

	if authDI.WithAuthInjector = authDI.NewAuthInjector(sharedDI.WithSharedInjector.DB); authDI.WithAuthInjector == nil {
		log.Println("could not initialize auth dependencies")
	}

	if walletDI.WithWalletInjector = walletDI.NewWalletInjector(sharedDI.WithSharedInjector.DB); walletDI.WithWalletInjector == nil {
		log.Println("could not initialize wallet dependencies")
	}
}
