package validator

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

type CustomValidator struct {
	Validator *validator.Validate
}

func NewValidator() *validator.Validate {
	return validator.New()
}

func (cv *CustomValidator) Validate(i interface{}) error {
	err := cv.Validator.Struct(i)
	if err != nil {
		errorMap := map[int]string{}
		for i, e := range err.(validator.ValidationErrors) {
			errorMap[i] = e.Field()
		}
		return errors.New(errorMap[0])
	}

	return nil
}
