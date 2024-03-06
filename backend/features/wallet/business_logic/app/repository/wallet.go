package repository

import (
	"go-wallet-api/features/wallet/business_logic/app/entities"
	"go-wallet-api/features/wallet/business_logic/app/models"
	"go-wallet-api/features/wallet/pkg"
	"gorm.io/gorm"
	"log"
)

const (
	GtWalletCountByOwnerIdQuery = "select count(*) as count from wallets where owner = ?"
	DeleteWalletByIdQuery       = "delete from wallets where id = ?"
)

type WalletRepository struct {
	pkg.IWalletRepository
	db *gorm.DB
}

func NewWalletRepository(db *gorm.DB) pkg.IWalletRepository {
	return &WalletRepository{db: db}
}

func (repo *WalletRepository) FindAllWallets(dataChan chan []*entities.WalletEntity, errChan chan error) {
	wallets := make([]*entities.WalletEntity, 0)
	if err := repo.db.Find(&wallets).Error; err != nil {
		log.Println(err)
		errChan <- err
		dataChan <- nil
		return
	}

	dataChan <- wallets
	errChan <- nil
}
func (repo *WalletRepository) FindWalletById(id string, dataChan chan *entities.WalletEntity, errChan chan error) {
	wallet := &entities.WalletEntity{}
	if err := repo.db.Where("id = ?", id).Find(&wallet).Error; err != nil {
		log.Println(err)
		errChan <- err
		dataChan <- nil
		return
	}
	dataChan <- wallet
	errChan <- nil
}
func (repo *WalletRepository) CreateNewWallet(wallet *entities.WalletEntity, errChan chan error) {
	if err := repo.db.Create(wallet).Error; err != nil {
		log.Println(err)
		errChan <- err
		return
	}

	errChan <- nil
}
func (repo *WalletRepository) UpdateWallet(wallet *entities.WalletEntity, errChan chan error) {
	if err := repo.db.Save(wallet).Error; err != nil {
		log.Println(err)
		errChan <- err
		return
	}

	errChan <- nil
}
func (repo *WalletRepository) GetWalletsCount(owner string, dataChan chan *int64, errChan chan error) {
	var count int64
	err := repo.db.Model(&entities.WalletEntity{}).Raw(GtWalletCountByOwnerIdQuery, owner).Count(&count).Error

	if err != nil {
		dataChan <- nil
		errChan <- err
		return
	}

	dataChan <- &count
	errChan <- nil
}

func (repo *WalletRepository) WalletExist(entity *entities.WalletEntity, dataChan chan bool, errChan chan error) {
	var wallet *models.Wallet

	err := repo.db.Where("account_number = ? and account_scheme = ?", entity.AccountNumber, entity.AccountScheme).Find(wallet).Error

	if err != nil {
		dataChan <- false
		errChan <- err
		return
	}

	errChan <- nil

	if wallet != nil {
		dataChan <- true
	} else {
		dataChan <- false

	}

}

func (repo *WalletRepository) DeleteWallet(id string, errChan chan error) {
	if err := repo.db.Exec(DeleteWalletByIdQuery, id).Error; err != nil {
		errChan <- err
		return
	}

	errChan <- nil
}
