package validator

import (
	"github.com/go-playground/validator/v10"
)

type Validator struct {
	validator *validator.Validate
}

func NewValidator() *Validator {
	return &Validator{
		validator: validator.New(),
	}
}

func (uv *Validator) Validate(entity interface{}) error {
	return uv.validator.Struct(entity)
}
