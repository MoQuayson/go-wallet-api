package services

import (
	"errors"
	"go-wallet-api/features/auth/business_logic/models"
	"go-wallet-api/features/auth/business_logic/requests"
	"go-wallet-api/features/auth/pkg"
	"go-wallet-api/features/shared/utils"
	"go-wallet-api/features/users/business_logic/entities"
	userModel "go-wallet-api/features/users/business_logic/models"
	usersPkg "go-wallet-api/features/users/pkg"
	"log"
)

var (
	errEmailExists = errors.New("email already exists")
)

type AuthService struct {
	pkg.IAuthService
	repo usersPkg.IUserRepository
}

func NewAuthService(repo usersPkg.IUserRepository) pkg.IAuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) AuthenticateUser(req *requests.LoginRequest) (*userModel.User, error) {
	userChan, errChan := utils.MakeDataAndErrorChannels[entities.UserEntity]()
	go s.repo.FindUserByEmail(req.Email, userChan, errChan)
	if err := <-errChan; err != nil {
		return nil, err
	}

	authUser := <-userChan

	if !s.isCorrectPassword(authUser.Password, req.Password) {
		return nil, nil
	}

	token, err := s.getJWTToken(authUser)
	if err != nil {
		return nil, err
	}
	authUser.Token = token

	return userModel.NewUserModelWithUserEntity(authUser), err
}

// SignUpUser signs up user
func (s *AuthService) SignUpUser(req *requests.SignUpRequest) error {
	userChan, userErrChan := utils.MakeDataAndErrorChannels[entities.UserEntity]()
	go s.repo.FindUserByEmail(req.Email, userChan, userErrChan)
	if err := <-userErrChan; err != nil {
		log.Printf("SignUpUser Err: %v", err)
		return err
	}

	if user := <-userChan; !user.ID.IsNil() {
		log.Printf("SignUpUser Err %v:", errEmailExists)
		return errEmailExists
	}

	errChan := make(chan error, 1)
	entity := models.NewUserEntityWithSignUpRequest(req)
	go s.repo.CreateNewUser(entity, errChan)

	if err := <-errChan; err != nil {
		log.Printf("SignUpUser Err: %v", err)
		return err
	}
	return nil
}
