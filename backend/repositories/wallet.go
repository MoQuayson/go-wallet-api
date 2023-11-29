package repositories

import (
	"fmt"
	"go-wallet-api/models"
	"go-wallet-api/requests"

	"strings"
	"sync"
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type WalletRepositoryInterface interface {
	FindAll() ([]models.Wallet, error)
	FindById(id string) (models.Wallet, error)
	Create(user models.Wallet) (models.Wallet, error)
	Update(id string, user models.Wallet) (models.Wallet, error)
	Exists(request requests.WalletRequest)
	Count() int64
}

type WalletRepository struct {
	DB        *gorm.DB
	WaitGroup *sync.WaitGroup
}

var walletRepository *WalletRepository

// Init wallet repository
func NewWalletRepository(db *gorm.DB) {
	walletRepository = &WalletRepository{DB: db, WaitGroup: &sync.WaitGroup{}}
}

// Get Repository
func GetWalletRepository() *WalletRepository {
	return walletRepository
}

// Gets all wallets from db
func (repo *WalletRepository) FindAll(channel chan models.DBResponse) {
	defer repo.WaitGroup.Done()
	wallets := []models.Wallet{}

	err := repo.DB.Preload("User").Find(&wallets).Error

	channel <- models.DBResponse{
		Data:  wallets,
		Error: err,
	}
}

// Get wallet by id from db
func (repo *WalletRepository) FindById(id string, channel chan models.DBResponse) {
	defer repo.WaitGroup.Done()
	wallet := models.Wallet{}

	err := repo.DB.Preload("User").Where("id = ?", id).First(&wallet).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		channel <- models.DBResponse{
			Data:  wallet,
			Error: err,
		}
	} else {
		channel <- models.DBResponse{
			Data:  wallet,
			Error: nil,
		}
	}
}

// gets wallet count
func (repo *WalletRepository) Count(userId any, channel chan models.DBResponse) {
	var count int64
	defer repo.WaitGroup.Done()
	err := repo.DB.Model(&models.Wallet{}).Where("owner = ?", userId).Count(&count).Error

	channel <- models.DBResponse{
		Data:  count,
		Error: err,
	}
}

// Check if wallet already exists
func (repo *WalletRepository) Exists(w requests.WalletRequest, channel chan models.DBResponse) {
	defer repo.WaitGroup.Done()

	var count int64
	err := repo.DB.Where("owner = ?", w.Owner).
		Where("account_number = ?", w.AccountNumber).
		Where("account_scheme = ?", w.AccountScheme).
		Count(&count).Error

	channel <- models.DBResponse{
		Data:  count >= 1,
		Error: err,
	}

}

// Add new wallet
func (repo *WalletRepository) Create(request requests.WalletRequest, channel chan models.DBResponse) {
	defer repo.WaitGroup.Done()
	w := models.Wallet{
		Name:          fmt.Sprintf("%s %s models.Wallet", request.AccountScheme, strings.ToUpper(request.Type)),
		Type:          request.Type,
		AccountNumber: TrimAccountNumber(request.Type, request.AccountNumber),
		AccountScheme: request.AccountScheme,
		Owner:         uuid.FromStringOrNil(request.Owner),
		CreatedAt:     time.Now(),
		User:          models.User{},
	}
	err := repo.DB.Create(&w).Error
	channel <- models.DBResponse{
		Data:  w,
		Error: err,
	}
}

// Updates wallet
func (repo *WalletRepository) Update(wallet models.Wallet, channel chan models.DBResponse) {
	defer repo.WaitGroup.Done()

	err := repo.DB.Model(&models.Wallet{}).Where("id = ?", wallet.ID).Updates(wallet).Error

	channel <- models.DBResponse{
		Data:  wallet,
		Error: err,
	}
}

// function to delete wallet from db
func (repo *WalletRepository) Delete(id string, channel chan models.DBResponse) {
	defer repo.WaitGroup.Done()

	err := repo.DB.Exec("delete from wallets where id = ?", id).Error

	channel <- models.DBResponse{
		Error: err,
	}
}
