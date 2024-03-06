package handlers

import (
	walletDI "go-wallet-api/features/wallet/di"
	"go-wallet-api/features/wallet/pkg"
)

const MaxWalletCount = 5

func getWalletService() pkg.IWalletService {
	return walletDI.WithWalletInjector.Service
}
