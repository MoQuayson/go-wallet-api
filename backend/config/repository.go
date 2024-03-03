package config

import (
	"go-wallet-api/repositories"

	"gorm.io/gorm"
)

/*
This is where you register your repositories
This is more like a DI style
*/
func RegisterRepositories(db *gorm.DB) {
	repositories.NewWalletRepository(db)
}
