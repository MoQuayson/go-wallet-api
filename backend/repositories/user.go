package repositories

import (
	"go-wallet-api/models"
	"go-wallet-api/requests"
	"sync"
	"time"

	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	FindAll(channel chan models.DBResponse)
	FindById(id string, channel chan models.DBResponse)
	Create(user models.User, channel chan models.DBResponse)
	Update(id string, channel chan models.DBResponse)
	Delete(id string, channel chan models.DBResponse)
	UserExistsByPhone(phone string, channel chan models.DBResponse)
}

type UserRepository struct {
	DB          *gorm.DB
	WaitGroup   *sync.WaitGroup
	IRepository UserRepositoryInterface
}

var userRepository *UserRepository

func NewUserRepository(db *gorm.DB) {
	userRepository = &UserRepository{DB: db, WaitGroup: &sync.WaitGroup{}}
}

func GetUserRepository() *UserRepository {
	return userRepository
}

// Finds all users in db
func (repo *UserRepository) FindAll(channel chan models.DBResponse) {
	defer repo.WaitGroup.Done()
	users := []models.User{}

	err := repo.DB.Find(&users).Error

	channel <- models.DBResponse{
		Data:  users,
		Error: err,
	}
}

// Find user by id
func (repo *UserRepository) FindById(id string, channel chan models.DBResponse) {
	defer repo.WaitGroup.Done()
	user := models.User{}
	err := repo.DB.Where("id = ?", id).First(&user).Error

	channel <- models.DBResponse{
		Data:  user,
		Error: err,
	}
}

// Inserts new user into db
func (repo *UserRepository) Create(req requests.UserRequest, channel chan models.DBResponse) {
	defer repo.WaitGroup.Done()
	u := models.User{
		Name:      req.Name,
		Email:     req.Email,
		PhoneNum:  models.ConvertToNullString(req.PhoneNum),
		Role:      "User",
		Password:  "password",
		CreatedAt: time.Now(),
	}

	//create user
	err := repo.DB.Create(&u).Error

	//write to channel
	channel <- models.DBResponse{
		Data:  u,
		Error: err,
	}
}

// Updates user
func (repo *UserRepository) Update(user models.User, channel chan models.DBResponse) {
	defer repo.WaitGroup.Done()
	err := repo.DB.Save(&user).Error

	//set data to channel
	channel <- models.DBResponse{
		Data:  user,
		Error: err,
	}

}

// function to check if user exists by email
func (repo *UserRepository) UserExistsByPhone(phone string, channel chan models.DBResponse) {
	defer repo.WaitGroup.Done()

	var count int64
	err := repo.DB.Model(models.User{}).Where("phone_num = ?", phone).Count(&count).Error

	//set channel data
	channel <- models.DBResponse{
		Data:  count >= 1,
		Error: err,
	}

}

// function to delete user from db
func (repo *UserRepository) Delete(id string, channel chan models.DBResponse) {
	defer repo.WaitGroup.Done()

	err := repo.DB.Exec("delete from users where id = ?", id).Error

	channel <- models.DBResponse{
		Error: err,
	}
}
