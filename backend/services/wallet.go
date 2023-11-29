package services

import (
	"go-wallet-api/models"
	"go-wallet-api/repositories"
	"go-wallet-api/requests"
	"sync"
)

type IWalletService interface {
	FindAll() ([]models.Wallet, error)
	FindById(id string) (models.Wallet, error)
	CreateNewWallet(request requests.WalletRequest) (models.Wallet, error)
	UpdateWallet(id string, request requests.WalletRequest) (models.Wallet, error)
	TrimAccountNum(walletType, account string) string
	WalletsCount(owner string) int64
}

type WalletService struct {
	Repository *repositories.WalletRepository
	WaitGroup  *sync.WaitGroup
}

var walletService *WalletService

// Init new service
func NewWalletService(walletRepo *repositories.WalletRepository) {
	walletService = &WalletService{Repository: walletRepo, WaitGroup: walletRepo.WaitGroup}
}

func GetWalletService() *WalletService {
	return walletService
}

// Find all wallets
func (service *WalletService) FindAll() ([]models.Wallet, error) {
	channel := make(chan models.DBResponse) //channel for db response

	wallets := []models.Wallet{}
	var err error

	service.WaitGroup.Add(1)
	go service.Repository.FindAll(channel)

	//wait for process to finish and close the channel
	WaitAndCloseChannel(service.WaitGroup, channel)

	//get data from channel
	for c := range channel {
		wallets = c.Data.([]models.Wallet)
		err = c.Error
	}

	return wallets, err
}

// Finds user by id
func (service *WalletService) FindById(id string) (models.Wallet, error) {
	channel := make(chan models.DBResponse) //channel for db response

	wallet := models.Wallet{}
	var err error

	service.WaitGroup.Add(1)
	go service.Repository.FindById(id, channel)

	//wait for process to finish and close the channel
	WaitAndCloseChannel(service.WaitGroup, channel)

	//get data from channel
	for c := range channel {
		wallet = c.Data.(models.Wallet)
		err = c.Error
	}

	return wallet, err
}

// Find all Wallets
func (service *WalletService) CreateNewWallet(request requests.WalletRequest) (models.Wallet, error) {
	channel := make(chan models.DBResponse) //channel for db response

	wallet := models.Wallet{}
	var err error

	service.WaitGroup.Add(1)
	go service.Repository.Create(request, channel)

	//wait for process to finish and close the channel
	WaitAndCloseChannel(service.WaitGroup, channel)

	//get data from channel
	for c := range channel {
		wallet = c.Data.(models.Wallet)
		err = c.Error
	}

	return wallet, err
}

// Find all Wallets
func (service *WalletService) UpdateWallet(wallet models.Wallet) (models.Wallet, error) {
	channel := make(chan models.DBResponse) //channel for db response

	var err error

	service.WaitGroup.Add(1)
	go service.Repository.Update(wallet, channel)

	//wait for process to finish and close the channel
	WaitAndCloseChannel(service.WaitGroup, channel)

	//get data from channel
	for c := range channel {
		wallet = c.Data.(models.Wallet)
		err = c.Error
	}

	return wallet, err
}

// services to delete user
func (service *WalletService) DeleteWallet(id string) error {
	var err error
	channel := make(chan models.DBResponse)

	service.WaitGroup.Add(1)
	go service.Repository.Delete(id, channel)

	//wait for process to finish and close the channel
	WaitAndCloseChannel(service.WaitGroup, channel)

	for c := range channel {
		err = c.Error
	}

	return err
}

func (service *WalletService) WalletsCount(owner string) int64 {
	channel := make(chan models.DBResponse) //channel for db response

	var count int64
	service.WaitGroup.Add(1)
	go service.Repository.Count(owner, channel)

	//wait for process to finish and close the channel
	WaitAndCloseChannel(service.WaitGroup, channel)

	//get data from channel
	for c := range channel {
		count = c.Data.(int64)
	}

	return count
}
func (service *WalletService) WalletExist(request requests.WalletRequest) bool {
	channel := make(chan models.DBResponse) //channel for db response

	var walletExist bool
	service.WaitGroup.Add(1)
	go service.Repository.Exists(request, channel)

	//wait for process to finish and close the channel
	WaitAndCloseChannel(service.WaitGroup, channel)

	//get data from channel
	for c := range channel {
		walletExist = c.Data.(bool)
	}

	return walletExist
}
