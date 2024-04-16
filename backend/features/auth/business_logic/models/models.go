package models

import (
	"go-wallet-api/features/auth/business_logic/requests"
	"go-wallet-api/features/shared/utils"
	"go-wallet-api/features/shared/utils/enums"
	"go-wallet-api/features/users/business_logic/entities"
)

func NewUserEntityWithSignUpRequest(req *requests.SignUpRequest) *entities.UserEntity {
	return &entities.UserEntity{
		ID:       utils.NewUUID(),
		Name:     req.Name,
		Email:    req.Email,
		PhoneNum: &req.PhoneNum,
		Role:     enums.RoleType_User,
		Password: req.Password,
	}
}
