package models

import (
	"fmt"
	"go-wallet-api/requests"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2/log"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

const (
	MAX_WALLET_COUNT int64  = 3
	MOMO_WALLET_TYPE string = "MOMO"
	CARD_WALLET_TYPE string = "CARD"
)

type Wallet struct {
	ID            uuid.UUID    `gorm:"column:id;type:uuid;primaryKey" json:"id" form:"id"`
	Name          string       `gorm:"column:name;size:255" json:"name" form:"name"`
	Type          string       `gorm:"column:type;size:255" json:"type" form:"type"`
	AccountNumber string       `gorm:"column:account_number;size:255;" json:"account_number" form:"account_number"`
	AccountScheme string       `gorm:"column:account_scheme;size:255" json:"account_scheme" form:"account_scheme"`
	Owner         uuid.UUID    `gorm:"column:owner;type:uuid" json:"owner" form:"owner"`
	User          User         `gorm:"foreignKey:Owner;references:ID" json:"user,omitempty"`
	CreatedAt     time.Time    `gorm:"column:created_at;type:timestamp" json:"created_at" form:"created_at"`
	UpdatedAt     NullDateTime `gorm:"column:updated_at;type:timestamp" json:"updated_at" form:"updated_at"`
}

func (Wallet) TableName() string {
	return "wallets"
}

func (w *Wallet) BeforeCreate(tx *gorm.DB) error {
	w.ID, _ = uuid.NewV4()
	w.CreatedAt = time.Now()

	return nil
}

// Gets all wallets from db
func GetAllWallets(db *gorm.DB) ([]Wallet, error) {
	wallets := []Wallet{}

	if err := db.Preload("User").Find(&wallets).Error; err != nil {
		return wallets, err
	}
	return wallets, nil
}

// Get wallet by id from db
func GetWalletById(id string, db *gorm.DB) (Wallet, error) {
	wallet := Wallet{}

	if err := db.Preload("User").Where("id = ?", id).First(&wallet).Error; err != nil && err != gorm.ErrRecordNotFound {
		return wallet, err
	}
	return wallet, nil
}

// gets wallet count
func GetUserWalletCount(userId any, db *gorm.DB) int64 {
	var count int64
	if err := db.Model(&Wallet{}).Where("owner = ?", userId).Count(&count).Error; err != nil {
		log.Errorf("Wallet: %s Error: %s GetUserWalletCount fn", userId, err)
	}

	log.Infof("Owner: %s Wallet Count: %v", userId, count)

	return count
}

// Check if wallet already exists
func UserWalletExists(w requests.WalletRequest, db *gorm.DB) int64 {
	var count int64
	if err := db.Where("owner = ?", w.Owner).
		Where("account_number = ?", w.AccountNumber).
		Where("account_scheme = ?", w.AccountScheme).
		Count(&count).Error; err != nil {
		log.Errorf("User: %s Error: %s at UserWalletExists fn", w.Owner, err)
	}

	return count
}

// Add new wallet
func CreateNewWallet(req requests.WalletRequest, db *gorm.DB) (Wallet, error) {
	w := Wallet{
		Name:          fmt.Sprintf("%s %s Wallet", req.AccountScheme, strings.ToUpper(req.Type)),
		Type:          req.Type,
		AccountNumber: TrimAccountNumber(req.Type, req.AccountNumber),
		AccountScheme: req.AccountScheme,
		Owner:         uuid.FromStringOrNil(req.Owner),
		CreatedAt:     time.Now(),
		User:          User{},
	}

	if err := db.Create(&w).Error; err != nil {
		log.Error(err)
		return w, err
	}

	log.Info("Wallet added")
	return w, nil
}

func TrimAccountNumber(walletType, account string) string {
	if strings.Contains(strings.ToLower(walletType), strings.ToLower(CARD_WALLET_TYPE)) {
		return account[:6]
	}

	return account
}

// Updates user
func UpdateWallet(id string, req requests.WalletRequest, db *gorm.DB) (Wallet, error) {
	wallet := Wallet{}
	var err error
	if err = db.Where("id = ?", id).First(&wallet).Error; err != nil {
		return wallet, err
	}

	wallet.Name = fmt.Sprintf("%s %s Wallet", req.AccountScheme, strings.ToUpper(req.Type))
	wallet.Type = req.Type
	wallet.AccountNumber = TrimAccountNumber(req.Type, req.AccountNumber)
	wallet.AccountScheme = req.AccountScheme

	err = db.Save(&wallet).Error
	if err != nil {
		return Wallet{}, err
	}

	return wallet, nil
}
