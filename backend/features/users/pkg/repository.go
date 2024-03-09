package pkg

import (
	"go-wallet-api/features/users/business_logic/app/entities"
)

type IUserRepository interface {
	FindAllUsers(chan []*entities.UserEntity, chan error)
	FindUserById(string, chan *entities.UserEntity, chan error)
	FindUserByEmail(string, chan *entities.UserEntity, chan error)
	CreateNewUser(*entities.UserEntity, chan error)
	UpdateUser(*entities.UserEntity, chan error)
	DeleteUser(string, chan error)
	CheckIfUserExistByPhone(string, chan *entities.UserEntity, chan error)
}
