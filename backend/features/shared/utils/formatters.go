package utils

import (
	"go-wallet-api/features/shared/utils/enums"
	"strings"
)

func TrimAccountNumberWithWalletType(accountNumber string, walletType enums.WalletType) string {
	if walletType == enums.WalletType_Card {
		return accountNumber[:6]
	}

	return accountNumber
}

func MapStringToWalletType(value string) enums.WalletType {
	switch strings.ToUpper(value) {
	case enums.WalletType_Card:
		return enums.WalletType_Card
	case enums.WalletType_Momo:
		return enums.WalletType_Momo
	default:
		return enums.WalletType_Invalid
	}
}
