package requests

type SignUpRequest struct {
	Name                 string `json:"name" form:"name" validate:"required,min=5,max=255"`
	Email                string `json:"email" form:"email" validate:"required,email"`
	PhoneNum             string `json:"phone_num" form:"phone_num" validate:"required,min=5,max=20"`
	Role                 string `json:"role" form:"role"`
	Password             string `json:"password" form:"password" validate:"required,min=5,max=50"`
	PasswordConfirmation string `json:"password_confirmation" form:"password_confirmation" validate:"required,confirm-password"`
}
