package pkg

import (
	"go-wallet-api/features/users/business_logic/models"
	"go-wallet-api/features/users/business_logic/requests"
)

type IUserService interface {
	FindAllUsers() ([]*models.User, error)
	FindUserById(string) (*models.User, error)
	CreateNewUser(*requests.UserRequest) (*models.User, error)
	UpdateUser(string, *requests.UserRequest) (*models.User, error)
	DeleteUser(string) error
	CheckIfUserExistByPhone(string) (*models.User, error)
}
