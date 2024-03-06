package services

import (
	"github.com/golang-jwt/jwt/v5"
	"go-wallet-api/features/auth/business_logic/app/models"
	"go-wallet-api/features/auth/pkg"
	"go-wallet-api/features/shared/utils"
	"go-wallet-api/features/users/business_logic/app/entities"
	userModel "go-wallet-api/features/users/business_logic/app/models"
	usersPkg "go-wallet-api/features/users/pkg"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

type AuthService struct {
	pkg.IAuthService
	repo usersPkg.IUserRepository
}

func NewAuthService(repo usersPkg.IUserRepository) pkg.IAuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) AuthenticateUser(req *models.LoginRequest) (*userModel.User, error) {
	userChan, errChan := utils.MakeDataAndErrorChannels[entities.UserEntity]()
	s.repo.FindUserByEmail(req.Email, userChan, errChan)
	if err := <-errChan; err != nil {
		return nil, err
	}

	authUser := <-userChan

	startTime := time.Now()
	if !s.checkPasswordHash(authUser.Password, req.Password) {
		return nil, nil
	}

	//generate token and claims policy
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = authUser.ID
	claims["name"] = authUser.Name
	claims["email"] = authUser.Email
	claims["role"] = authUser.Role
	claims["phone_num"] = authUser.PhoneNum
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(utils.GetJWTSecret()))
	if err != nil {
		return nil, err
	}
	authUser.Token = t
	endTime := time.Now()
	difference := endTime.Sub(startTime)
	log.Printf("Seconds: %d ms\n", difference.Milliseconds())

	return userModel.NewUserModelWithUserEntity(authUser), err
}

func (s *AuthService) checkPasswordHash(hashPassword, password string) bool {
	var err error
	go func(error, string, string) {
		err = bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	}(err, hashPassword, password)
	return err == nil
}
