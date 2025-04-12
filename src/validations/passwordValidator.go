package validations

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)


func PasswoordValidator(fld validator.FieldLevel) bool {
	value, ok := fld.Field().Interface().(string)
	if !ok {
		return false
	}
	res, err := regexp.MatchString(`^[@#$&a-zA-Z0-9]{5,25}$`, value)
	if err != nil {
		panic(err)
	}
	return res
}