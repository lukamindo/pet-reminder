package validator

import (
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

// Validate validates
func ValidateStruct(i interface{}) error {
	if validate == nil {
		validate = validator.New()
	}
	return validate.Struct(i)
}
