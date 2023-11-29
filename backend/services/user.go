package services

import (
	"go-wallet-api/models"
	"go-wallet-api/repositories"
	"go-wallet-api/requests"
	"sync"
)

type IUserService interface {
	FindAll() ([]models.User, error)
	FindById(id string) (models.User, error)
	CreateNewUser(request requests.UserRequest) (models.User, error)
	UpdateUser(id string, request requests.UserRequest) (models.User, error)
}

type UserService struct {
	Repository *repositories.UserRepository
	WaitGroup  *sync.WaitGroup
}

var userService *UserService

// Init new service
func NewUserService(repository *repositories.UserRepository) {
	userService = &UserService{Repository: repository, WaitGroup: repository.WaitGroup}
}
func GetUserService() *UserService {
	return userService
}

// Find all Users
func (service *UserService) FindAll() ([]models.User, error) {
	channel := make(chan models.DBResponse) //channel for db response

	users := []models.User{}
	var err error

	service.WaitGroup.Add(1)
	go service.Repository.FindAll(channel)

	//wait for process to finish and close the channel
	WaitAndCloseChannel(service.WaitGroup, channel)

	//get data from channel
	for c := range channel {
		users = c.Data.([]models.User)
		err = c.Error
	}

	return users, err
}

// Finds user by id
func (service *UserService) FindById(id string) (models.User, error) {
	user := models.User{}
	var err error
	channel := make(chan models.DBResponse)

	service.WaitGroup.Add(1)
	go service.Repository.FindById(id, channel)

	//wait for process to finish and close the channel
	WaitAndCloseChannel(service.WaitGroup, channel)

	for c := range channel {
		user = c.Data.(models.User) // cast
		err = c.Error
	}

	return user, err
}

// Create new user
func (service *UserService) CreateNewUser(request requests.UserRequest) (models.User, error) {
	user := models.User{}
	var err error
	channel := make(chan models.DBResponse)

	service.WaitGroup.Add(1)
	go service.Repository.Create(request, channel)

	//wait for process to finish and close the channel
	WaitAndCloseChannel(service.WaitGroup, channel)

	for c := range channel {
		user = c.Data.(models.User)
		err = c.Error
	}

	return user, err
}

// Update user
func (service *UserService) UpdateUser(id string, user models.User) (models.User, error) {
	var err error
	channel := make(chan models.DBResponse)

	service.WaitGroup.Add(1)
	go service.Repository.Update(user, channel)

	//wait for process to finish and close the channel
	WaitAndCloseChannel(service.WaitGroup, channel)

	for c := range channel {
		user = c.Data.(models.User)
		err = c.Error
	}

	return user, err

}

// services to delete user
func (service *UserService) DeleteUser(id string) error {
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

// Function to check if user exists
func (service *UserService) UserExistsByPhone(phone string) (bool, error) {
	var err error
	channel := make(chan models.DBResponse)
	var userExists bool

	service.WaitGroup.Add(1)
	go service.Repository.UserExistsByPhone(phone, channel)

	//wait for process to finish and close the channel
	WaitAndCloseChannel(service.WaitGroup, channel)

	for c := range channel {
		userExists = c.Data.(bool)
		err = c.Error
	}

	return userExists, err

}
