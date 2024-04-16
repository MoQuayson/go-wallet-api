package utils

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"go-wallet-api/features/auth/business_logic/requests"
	"reflect"
)

func NewValidation() *validator.Validate {
	v := validator.New()
	v.RegisterValidation("is-unique", ValidateIsUnique)
	v.RegisterValidation("confirm-password", ValidatePasswordConfirmation)

	return v
}

func ValidatePasswordConfirmation(fl validator.FieldLevel) bool {
	parent := fl.Parent()
	val := reflect.ValueOf(parent.Interface()).Type()
	if val == reflect.ValueOf(requests.SignUpRequest{}).Type() {
		req := parent.Interface().(requests.SignUpRequest)

		return req.Password == req.PasswordConfirmation
	}
	return true
}

func ValidateIsUnique(fl validator.FieldLevel) bool {
	return false
}

func GetValidationErrors(err error, model interface{}) map[string]string {
	//field := map[string][]string{}
	fieldMap := make(map[string]string, 0)
	val := reflect.ValueOf(model)

	for i := 0; i < val.Type().NumField(); i++ {
		for _, err := range err.(validator.ValidationErrors) {

			if val.Type().Field(i).Name == err.Field() {
				field := val.Type().Field(i).Tag.Get("json")
				fieldMap[field] = getValidationErrorMessage(field, err)
			}

		}
	}

	return fieldMap
}

func getValidationErrorMessage(field string, fieldError validator.FieldError) string {
	switch fieldError.Tag() {
	case "required":
		return fmt.Sprintf("The %s field is required", field)
	case "email":
		return fmt.Sprintf("The %s field must be a valid email address", field)
	case "min":
		return fmt.Sprintf("The %s field must be minimum of %s characters", field, fieldError.Param())
	case "max":
		return fmt.Sprintf("The %s field must be maximum of %s characters", field, fieldError.Param())
	case "confirm-password":
		return "The password and password_confirmation fields do not match"
	}
	return fieldError.Error()
}
