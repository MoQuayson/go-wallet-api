package validations

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"reflect"
)

var Validator *validator.Validate

func InitValidation() {
	Validator = validator.New()
	Validator.RegisterValidation("is-unique", ValidateIsUnique)
}

func ValidateIsUnique(fl validator.FieldLevel) bool {

	return false
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
	}
	return fieldError.Error()
}

func GetValidationErrors(err error, model interface{}) map[string][]string {
	//field := map[string][]string{}
	fieldMap := make(map[string][]string, 0)
	val := reflect.ValueOf(model)

	for i := 0; i < val.Type().NumField(); i++ {
		for _, err := range err.(validator.ValidationErrors) {

			if val.Type().Field(i).Name == err.Field() {
				field := val.Type().Field(i).Tag.Get("json")
				fieldMap[field] = append(fieldMap[field], getValidationErrorMessage(field, err))
			}

		}
	}

	return fieldMap
}
