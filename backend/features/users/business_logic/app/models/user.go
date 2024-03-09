package models

import (
	"github.com/gofrs/uuid"
	"go-wallet-api/features/shared/utils"
	"go-wallet-api/features/shared/utils/enums"
	"go-wallet-api/features/users/business_logic/app/entities"
	"time"
)

type UserRequest struct {
	Name     string `json:"name" form:"name" validate:"required,min=5,max=255"`
	Email    string `json:"email" form:"email" validate:"required,email"`
	PhoneNum string `json:"phone_num" form:"phone_num" validate:"required,min=5,max=20"`
	Role     string `json:"role" form:"role"`
}

type User struct {
	ID        uuid.UUID  `json:"id" form:"id"`
	Name      string     `json:"name" form:"name"`
	Email     string     `json:"email" form:"email"`
	PhoneNum  *string    `json:"phone_num" form:"phone_num"`
	Role      string     `json:"role" form:"role"`
	Password  string     `json:"-" form:"-"`
	Token     string     `json:"token,omitempty" form:"token,omitempty"`
	CreatedAt time.Time  `json:"created_at,omitempty" form:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" form:"updated_at"`
}

func NewUserModelWithUserEntity(user *entities.UserEntity) *User {
	return &User{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		PhoneNum:  user.PhoneNum,
		Role:      user.Role,
		Password:  user.Password,
		Token:     user.Token,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func NewUserEntity(req *UserRequest) *entities.UserEntity {
	if len(req.Role) == 0 {
		req.Role = enums.RoleType_User
	}
	return &entities.UserEntity{
		ID:       utils.NewUUID(),
		Name:     req.Name,
		Email:    req.Email,
		PhoneNum: &req.PhoneNum,
		Role:     req.Role,
	}
}
