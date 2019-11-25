package _interface

import "github.com/go-playground/validator"

type ValidatorInterface interface {
	Validation(fieldLevel validator.FieldLevel)
}
