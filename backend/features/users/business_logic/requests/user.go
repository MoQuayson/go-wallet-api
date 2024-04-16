package requests

type UserRequest struct {
	Name     string `json:"name" form:"name" validate:"required,min=5,max=255"`
	Email    string `json:"email" form:"email" validate:"required,email"`
	PhoneNum string `json:"phone_num" form:"phone_num" validate:"required,min=5,max=20"`
	Role     string `json:"role" form:"role"`
}
