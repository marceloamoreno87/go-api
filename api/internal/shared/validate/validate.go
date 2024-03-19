package validate

import (
	"github.com/go-playground/validator/v10"
)

type Validator struct {
	Data any
}

func NewValidator(data any) *Validator {
	return &Validator{
		Data: data,
	}
}

func (v *Validator) Validate() (err error) {
	validate := validator.New()
	err = validate.Struct(v.Data)
	return
}
