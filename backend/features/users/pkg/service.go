package pkg

import (
	"go-wallet-api/features/users/business_logic/app/models"
)

type IUserService interface {
	FindAllUsers() ([]*models.User, error)
	FindUserById(string) (*models.User, error)
	CreateNewUser(*models.UserRequest) (*models.User, error)
	UpdateUser(string, *models.UserRequest) (*models.User, error)
	DeleteUser(string) error
	CheckIfUserExistByPhone(string) (*models.User, error)
}
