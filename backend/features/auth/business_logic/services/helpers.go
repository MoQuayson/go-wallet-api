package services

import (
	"github.com/golang-jwt/jwt/v5"
	"go-wallet-api/features/shared/utils"
	userEnt "go-wallet-api/features/users/business_logic/entities"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

func (s *AuthService) isCorrectPassword(hashPassword, password string) bool {
	errChan := make(chan error, 1)
	go func(chan error, string, string) {
		err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
		if err != nil {
			log.Printf("isCorrectPassword: %v", err)
			errChan <- err
			return
		}

		errChan <- nil
	}(errChan, hashPassword, password)

	err := <-errChan
	return err == nil
}

func (s *AuthService) getJWTToken(authUser *userEnt.UserEntity) (string, error) {
	//generate token and claims policy
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = authUser.ID
	claims["name"] = authUser.Name
	claims["email"] = authUser.Email
	claims["role"] = authUser.Role
	claims["phone_num"] = authUser.PhoneNum
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	return token.SignedString([]byte(utils.GetJWTSecret()))
}
