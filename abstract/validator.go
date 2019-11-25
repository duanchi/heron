package abstract

import (
	"github.com/go-playground/validator"
	_interface "go.heurd.com/heron-go/heron/interface"
)

type Validator struct {
	Bean
	_interface.ValidatorInterface
}

func (this *Validator) Validation (fieldLevel validator.FieldLevel) bool {
	return true
}
