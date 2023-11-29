package repositories

import (
	"go-wallet-api/models"
	"strings"
)

func TrimAccountNumber(walletType, account string) string {
	if strings.Contains(strings.ToLower(walletType), strings.ToLower(models.CARD_WALLET_TYPE)) {
		return account[:6]
	}

	return account
}
