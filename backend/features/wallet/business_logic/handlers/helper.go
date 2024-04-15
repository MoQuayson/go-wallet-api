package handlers

import (
	"fmt"
	"go-wallet-api/features/shared/utils/enums"
)

const (
	MaxWalletCount = 5
)

var (
	MaxWalletCountMsg = enums.ResponseMsg(fmt.Sprintf("Cannot have more than %v wallets", MaxWalletCount))
)
