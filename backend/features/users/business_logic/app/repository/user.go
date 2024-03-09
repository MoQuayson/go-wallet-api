package repository

import (
	"go-wallet-api/features/users/business_logic/app/entities"
	"go-wallet-api/features/users/pkg"
	"gorm.io/gorm"
)

const (
	DeleteUserByIdQuery = "delete from users where id = ?"
	GetUserByPhoneQuery = "where phone = ?"
)

type UserRepository struct {
	pkg.IUserRepository
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) pkg.IUserRepository {
	return &UserRepository{db: db}
}

// FindAllUsers retrieves all users from db
func (repo *UserRepository) FindAllUsers(dataChan chan []*entities.UserEntity, errChan chan error) {
	users := make([]*entities.UserEntity, 0)

	if err := repo.db.Find(&users).Error; err != nil {
		dataChan <- nil
		errChan <- err
		return
	}

	dataChan <- users
	errChan <- nil
}

// FindUserById gets user by id
func (repo *UserRepository) FindUserById(id string, dataChan chan *entities.UserEntity, errChan chan error) {
	user := &entities.UserEntity{}

	if err := repo.db.Where("id = ?", id).Find(&user).Error; err != nil {
		dataChan <- nil
		errChan <- err
		return
	}

	dataChan <- user
	errChan <- nil
}
func (repo *UserRepository) FindUserByEmail(email string, dataChan chan *entities.UserEntity, errChan chan error) {
	user := &entities.UserEntity{}

	if err := repo.db.Where("email = ?", email).Find(&user).Error; err != nil {
		dataChan <- nil
		errChan <- err
		return
	}

	dataChan <- user
	errChan <- nil
}
func (repo *UserRepository) CreateNewUser(user *entities.UserEntity, errChan chan error) {

	if err := repo.db.Create(user).Error; err != nil {
		errChan <- err
		return
	}

	errChan <- nil
}
func (repo *UserRepository) UpdateUser(user *entities.UserEntity, errChan chan error) {
	if err := repo.db.Save(&user).Error; err != nil {
		errChan <- err
		return
	}

	errChan <- nil
}
func (repo *UserRepository) DeleteUser(id string, errChan chan error) {
	if err := repo.db.Exec("delete from users where id = ?", id).Error; err != nil {
		errChan <- err
		return
	}

	errChan <- nil
}
func (repo *UserRepository) CheckIfUserExistByPhone(phoneNum string, dataChan chan *entities.UserEntity, errChan chan error) {
	user := &entities.UserEntity{}

	if err := repo.db.Where(GetUserByPhoneQuery, phoneNum).Find(&user).Error; err != nil {
		dataChan <- nil
		errChan <- err
		return
	}

	dataChan <- user
	errChan <- nil
}
