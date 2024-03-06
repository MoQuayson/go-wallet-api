package models

import (
	"github.com/gofrs/uuid"
	"go-wallet-api/features/shared/utils"
	userModel "go-wallet-api/features/users/business_logic/app/models"
	"go-wallet-api/features/wallet/business_logic/app/entities"
	"time"
)

type Wallet struct {
	ID            uuid.UUID       `json:"id" form:"id"`
	Name          string          `json:"name" form:"name"`
	Type          string          `json:"type" form:"type"`
	AccountNumber string          `json:"account_number" form:"account_number"`
	AccountScheme string          `json:"account_scheme" form:"account_scheme"`
	Owner         uuid.UUID       `json:"owner" form:"owner"`
	User          *userModel.User `json:"user,omitempty"`
	CreatedAt     time.Time       `json:"created_at" form:"created_at"`
	UpdatedAt     *time.Time      `json:"updated_at" form:"updated_at"`
}

func NewWalletModelWithWalletEntity(wallet *entities.WalletEntity) *Wallet {
	if wallet.ID.IsNil() {
		return nil
	}

	return &Wallet{
		ID:            wallet.ID,
		Name:          wallet.Name,
		Type:          wallet.Type,
		AccountNumber: wallet.AccountNumber,
		AccountScheme: wallet.AccountScheme,
		Owner:         wallet.Owner,
		CreatedAt:     wallet.CreatedAt,
		UpdatedAt:     wallet.UpdatedAt,
	}
}

func NewWalletEntity(req *WalletRequest) *entities.WalletEntity {
	owner, _ := uuid.FromString(req.Owner)
	accountNumber := utils.TrimAccountNumberWithWalletType(req.AccountNumber, utils.MapStringToWalletType(req.Type))
	return &entities.WalletEntity{
		ID:            utils.NewUUID(),
		Name:          utils.GenerateWalletName(req.AccountScheme, req.Type),
		Type:          req.Type,
		AccountNumber: accountNumber,
		AccountScheme: req.AccountScheme,
		Owner:         owner,
		CreatedAt:     time.Now(),
	}
}
