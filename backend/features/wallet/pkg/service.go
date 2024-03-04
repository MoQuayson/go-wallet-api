package pkg

import (
	"go-wallet-api/features/wallet/business_logic/app/models"
)

type IWalletService interface {
	FindAllWallets() ([]*models.Wallet, error)
	FindWalletById(string) (*models.Wallet, error)
	CreateNewWallet(*models.WalletRequest) (*models.Wallet, error)
	UpdateWallet(string, models.WalletRequest) (*models.Wallet, error)
	GetWalletsCount(string) (*int64, error)
}
