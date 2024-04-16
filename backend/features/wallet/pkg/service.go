package pkg

import (
	models2 "go-wallet-api/features/wallet/business_logic/models"
)

type IWalletService interface {
	FindAllWallets() ([]*models2.Wallet, error)
	FindWalletById(string) (*models2.Wallet, error)
	CreateNewWallet(*models2.WalletRequest) (*models2.Wallet, error)
	UpdateWallet(string, *models2.WalletRequest) (*models2.Wallet, error)
	GetWalletsCount(string) (*int64, error)
	DeleteWallet(string) error
}
