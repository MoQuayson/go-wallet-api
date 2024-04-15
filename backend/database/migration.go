package database

import (
	"go-wallet-api/features/shared/di"
	userEntity "go-wallet-api/features/users/business_logic/entities"
	walletEntity "go-wallet-api/features/wallet/business_logic/entities"
	"log"
)

// MigrateEntities Migrate Entities Here
func MigrateEntities() {
	db := di.WithSharedInjector.DB
	err := db.AutoMigrate(
		&userEntity.UserEntity{},
		&walletEntity.WalletEntity{},
	)

	if err != nil {
		log.Fatalln("failed to migrate entities")
	}

}
