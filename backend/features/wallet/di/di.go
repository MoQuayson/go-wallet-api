package di

import (
	"go-wallet-api/features/wallet/business_logic/app/repository"
	"go-wallet-api/features/wallet/business_logic/services"
	"go-wallet-api/features/wallet/pkg"
	"gorm.io/gorm"
)

var WithWalletInjector *WalletInjector

type WalletInjector struct {
	Service pkg.IWalletService
	db      *gorm.DB
}

func NewWalletInjector(db *gorm.DB) *WalletInjector {
	injector := &WalletInjector{}
	injector.Service = services.NewWalletService(repository.NewWalletRepository(db))
	return injector
}
