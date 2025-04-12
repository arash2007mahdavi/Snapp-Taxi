package validations

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	Property string `json:"Property"`
	Tag      string `json:"Tag"`
	Value    string `json:"Value"`
}

func GenerateValidationError(err error) *[]ValidationError {
	var ValidationErrors []ValidationError
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		for _, err := range err.(validator.ValidationErrors) {
			var el ValidationError
			el.Property = err.Field()
			el.Tag = err.Tag()
			el.Value = err.Param()
			ValidationErrors = append(ValidationErrors, el)
		}
		return &ValidationErrors
	}
	return nil
}
