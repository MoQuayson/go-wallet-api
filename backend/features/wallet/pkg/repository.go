package pkg

import (
	"go-wallet-api/features/wallet/business_logic/entities"
)

type IWalletRepository interface {
	FindAllWallets(chan []*entities.WalletEntity, chan error)
	FindWalletById(string, chan *entities.WalletEntity, chan error)
	CreateNewWallet(*entities.WalletEntity, chan error)
	UpdateWallet(*entities.WalletEntity, chan error)
	GetWalletsCount(string, chan *int64, chan error)
	WalletExist(*entities.WalletEntity, chan bool, chan error)
	DeleteWallet(string, chan error)
}
