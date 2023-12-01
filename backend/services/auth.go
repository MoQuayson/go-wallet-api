package services

import (
	"go-wallet-api/models"
	"go-wallet-api/repositories"
	"go-wallet-api/requests"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type IAuthService interface {
	AuthenticateUser(requests.LoginRequest) (bool, models.User)
}

type AuthService struct {
	Repository *repositories.UserRepository
	WaitGroup  *sync.WaitGroup
}

var authService *AuthService

// Init new service
func NewAuthService(repository *repositories.UserRepository) {
	authService = &AuthService{Repository: repository, WaitGroup: repository.WaitGroup}
}

// Gets authservice
func GetAuthService() *AuthService {
	return authService
}

// This functions authenticates user credentials
func (s *AuthService) AuthenticateUser(request requests.LoginRequest) (bool, models.User) {
	channel := make(chan models.DBResponse)
	user := models.User{}
	var err error
	s.WaitGroup.Add(1)
	go s.Repository.FindByEmail(request.Email, channel)

	WaitAndCloseChannel(s.WaitGroup, channel)

	for c := range channel {
		user = c.Data.(models.User)
		err = c.Error
	}

	if err != nil || !checkPasswordHash(user.Password, request.Password) {
		log.Error(err)
		return false, user
	}

	//genereate token and claims policy
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["name"] = user.Name
	claims["email"] = user.Email
	claims["role"] = user.Role
	claims["phone_num"] = user.PhoneNum.String
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(GetJWTSecret()))
	if err != nil {
		return false, user
	}
	user.Token = t

	return true, user
}

func checkPasswordHash(hashPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	return err == nil
}
