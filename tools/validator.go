package tools

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

func Validate(data interface{}) error {
	validate := validator.New()
	if err := validate.Struct(data); err != nil {
		return errors.New("validation error")
	}

	return nil
}
