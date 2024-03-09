package services

import (
	"github.com/gofrs/uuid"
	"go-wallet-api/features/shared/utils"
	"go-wallet-api/features/wallet/business_logic/app/entities"
	"go-wallet-api/features/wallet/business_logic/app/models"
	"go-wallet-api/features/wallet/pkg"
	"time"
)

type WalletService struct {
	pkg.IWalletService
	repo pkg.IWalletRepository
}

func NewWalletService(repo pkg.IWalletRepository) pkg.IWalletService {
	return &WalletService{repo: repo}
}

func (s *WalletService) FindAllWallets() ([]*models.Wallet, error) {
	walletsChan, errChan := utils.MakeDataSliceAndErrorChannels[entities.WalletEntity]()
	go s.repo.FindAllWallets(walletsChan, errChan)
	walletEntities := <-walletsChan
	err := <-errChan

	if err != nil {
		return nil, err
	}

	wallets := make([]*models.Wallet, 0)
	for _, user := range walletEntities {
		wallets = append(wallets, models.NewWalletModelWithWalletEntity(user))
	}

	return wallets, nil
}
func (s *WalletService) FindWalletById(id string) (*models.Wallet, error) {
	walletChan, errChan := utils.MakeDataAndErrorChannels[entities.WalletEntity]()
	go s.repo.FindWalletById(id, walletChan, errChan)
	wallet := <-walletChan
	err := <-errChan

	if err != nil {
		return nil, err
	}

	return models.NewWalletModelWithWalletEntity(wallet), nil
}
func (s *WalletService) CreateNewWallet(req *models.WalletRequest) (*models.Wallet, error) {
	errChan := make(chan error, 1)
	wallet := models.NewWalletEntity(req)
	go s.repo.CreateNewWallet(wallet, errChan)
	err := <-errChan

	if err != nil {
		return nil, err
	}

	return models.NewWalletModelWithWalletEntity(wallet), nil
}
func (s *WalletService) UpdateWallet(id string, req *models.WalletRequest) (*models.Wallet, error) {
	//get wallet by id
	walletChan, errChan := utils.MakeDataAndErrorChannels[entities.WalletEntity]()
	go s.repo.FindWalletById(id, walletChan, errChan)
	walletEntity := <-walletChan
	err := <-errChan
	if err != nil {
		return nil, err
	}

	accountNumber := utils.TrimAccountNumberWithWalletType(req.AccountNumber, utils.MapStringToWalletType(req.Type))
	updatedAt := time.Now()
	walletEntity.Name = utils.GenerateWalletName(req.AccountScheme, req.Type)
	walletEntity.Type = req.Type
	walletEntity.AccountNumber = accountNumber
	walletEntity.AccountScheme = req.AccountScheme
	walletEntity.Owner = uuid.FromStringOrNil(req.Owner)
	walletEntity.UpdatedAt = &updatedAt

	errChan = make(chan error, 1)
	go s.repo.UpdateWallet(walletEntity, errChan)
	if err = <-errChan; err != nil {
		return nil, err
	}

	return models.NewWalletModelWithWalletEntity(walletEntity), nil
}
func (s *WalletService) GetWalletsCount(id string) (*int64, error) {
	walletCountChan, errChan := utils.MakeDataAndErrorChannels[int64]()
	go s.repo.GetWalletsCount(id, walletCountChan, errChan)
	walletCount := <-walletCountChan
	err := <-errChan

	if err != nil {
		return nil, err
	}

	return walletCount, nil
}

func (s *WalletService) DeleteWallet(id string) error {
	errChan := make(chan error, 1)
	s.repo.DeleteWallet(id, errChan)

	return <-errChan
}
